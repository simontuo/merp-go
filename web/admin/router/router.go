package router

import (
	"github.com/gin-gonic/gin"
	"github.com/simontuo/merp-go/config"
	"github.com/simontuo/merp-go/helper/err"
	"github.com/simontuo/merp-go/plugin/tracer/web"
	"github.com/simontuo/merp-go/web/admin/handler"
	. "github.com/simontuo/merp-go/web/admin/middleware"
)

const version = "v1"

func Init() *gin.Engine {
	router := gin.Default()
	// 注册自定义验证器
	RegisterValidator()
	// 404
	router.NoMethod(err.HandleMethodNotAllowed)
	router.NoRoute(err.HandleNotFound)

	tokenHandler := new(handler.TokenHandler)
	userHandler := new(handler.UserHandler)
	departmentHandler := new(handler.DepartmentHandler)
	tenantHandler := new(handler.TenantHandler)
	roleHandler := new(handler.RoleHandler)
	permissionHandler := new(handler.PermissionHandler)
	serviceHandler := new(handler.ServiceHandler)
	smsHandler := new(handler.SmsHandler)
	// 地址前缀
	var v *gin.RouterGroup
	if config.G_Config.Dev {
		v = router.Group("/admin/dev/api/v1")
	} else {
		v = router.Group("/admin/prod/api/v1")
	}

	v.Use(web.TracerWrapper)
	v.POST("/reset_password", Wrapper(userHandler.ResetPassword))
	// batch
	batch := v.Group("/batch")
	{
		batch.PUT("/users/ban", userHandler.BatchBan())
		batch.PUT("/departments/ban", departmentHandler.BatchBan())
		batch.PUT("/permissions/ban", permissionHandler.BatchBan())
		batch.PUT("/roles/ban", roleHandler.BatchBan())
	}
	// list
	list := v.Group("/list")
	{
		list.GET("/users", userHandler.List())
		list.GET("/admin_users", userHandler.AdminUserList())
	}
	// tree
	tree := v.Group("/trees")
	{
		tree.GET("/departments", departmentHandler.Tree())
		tree.GET("/permissions", permissionHandler.Tree())
	}
	// token handler
	token := v.Group("/token")
	{
		token.POST("/sign", Wrapper(tokenHandler.Sign))
		token.GET("/destroy", Wrapper(tokenHandler.Destroy))
	}
	// user handler
	user := v.Group("/users")
	{
		user.GET("/:id", userHandler.Info())
		user.GET("", userHandler.Page())
		user.POST("", userHandler.Store())
		user.POST("/bind/admin", userHandler.BindAdmin())
	}
	// department handler
	department := v.Group("/departments")
	{
		department.POST("", departmentHandler.Store())
		department.GET("/:id", departmentHandler.Info())
		department.PUT("/:id", departmentHandler.Update())
	}
	// tenant handler
	tenant := v.Group("/tenants")
	{
		tenant.GET("", tenantHandler.Page())
		tenant.POST("", tenantHandler.Store())
		tenant.GET("/:id", tenantHandler.Info())
		tenant.PUT("/:id", tenantHandler.Update())
	}
	// role handler
	role := v.Group("/roles")
	{
		role.GET("", roleHandler.Page())
		role.GET("/:id", roleHandler.Info())
		role.POST("", roleHandler.Store())
		role.PUT("/:id", roleHandler.Update())
		role.PUT("/:id/bind/users", roleHandler.BindUser())
	}
	// permission handler
	permission := v.Group("/permissions")
	{
		permission.GET("", permissionHandler.Page())
		permission.GET("/:id", permissionHandler.Info())
		permission.POST("", permissionHandler.Store())
		permission.PUT("/:id", permissionHandler.Update())
	}
	// service handler
	service := v.Group("/services")
	{
		service.GET("", serviceHandler.List())
		service.GET("/:name", serviceHandler.Info())
	}
	// sms handler
	sms := v.Group("/sms")
	{
		sms.GET("/login_verify_code", Wrapper(smsHandler.LoginVerifyCode))
		sms.GET("/reset_password_verify_code", Wrapper(smsHandler.ForgetPasswordVerifyCode))
	}
	return router
}
