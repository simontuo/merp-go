package handler

import (
	"net/http"
	"strconv"

	"github.com/simontuo/merp-go/helper/rsp"
	"github.com/simontuo/merp-go/web/admin/request/user"

	"github.com/gomodule/redigo/redis"
	apiErr "github.com/simontuo/merp-go/helper/err"
	redis2 "github.com/simontuo/merp-go/lib/redis"

	"github.com/simontuo/merp-go/config"
	"github.com/simontuo/merp-go/plugin/tracer/web"

	"github.com/gin-gonic/gin"
	"github.com/simontuo/merp-go/helper/response"
	srv_user "github.com/simontuo/merp-go/srv/userservice/proto/user"
)

type UserHandler struct {
	UserClient srv_user.UserService
}

func (h *UserHandler) cli() *UserHandler {
	h.UserClient = srv_user.NewUserService("go.micro.srv.user", config.WebService.Options().Service.Client())
	return h
}

// 获取用户信息
func (h *UserHandler) Info() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 链路追踪上下文
		ctx, ok := web.ContextWithSpan(c)
		if !ok {
			response.Err(c, http.StatusInternalServerError, "上下文出错")
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response.Err(c, http.StatusNotFound, "参数有误")
			return
		}
		req := srv_user.GetUserInfoRequest{
			Id: int32(id),
		}
		rsp, err := h.cli().UserClient.GetUserInfo(ctx, &req)
		if err != nil {
			response.MicroErr(c, err)
			return
		}

		response.OK(c, rsp.Data)
	}
}

// 获取用户信息
func (h *UserHandler) Page() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 链路追踪上下文
		ctx, ok := web.ContextWithSpan(c)
		if !ok {
			response.Err(c, http.StatusInternalServerError, "上下文出错")
			return
		}

		req := srv_user.GetUserPageRequest{}
		if err := c.ShouldBindQuery(&req); err != nil {
			response.Err(c, http.StatusBadRequest, "参数有误")
			return
		}

		res, err := h.cli().UserClient.GetUserPage(ctx, &req)
		if err != nil {
			response.MicroErr(c, err)
			return
		}

		data := make(map[string]interface{})
		data["items"] = res.Data
		data["total"] = res.Total

		response.OK(c, data)
	}
}

// 新增用户
func (h *UserHandler) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 链路追踪上下文
		ctx, ok := web.ContextWithSpan(c)
		if !ok {
			response.Err(c, http.StatusInternalServerError, "上下文出错")
			return
		}

		req := srv_user.StoreUserRequest{}
		if err := c.ShouldBindJSON(&req.Form); err != nil {
			response.Err(c, http.StatusBadRequest, "参数有误")
			return
		}

		// 新增用户
		res, err := h.cli().UserClient.StoreUser(ctx, &req)
		if err != nil {
			response.MicroErr(c, err)
			return
		}

		response.MsgOK(c, res.Data, "新增成功")
	}
}

// 批量禁用用户
func (h *UserHandler) BatchBan() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 链路追踪上下文
		ctx, ok := web.ContextWithSpan(c)
		if !ok {
			response.Err(c, http.StatusInternalServerError, "上下文出错")
			return
		}

		var req srv_user.BatchBanUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			response.Err(c, http.StatusBadRequest, "参数有误")
			return
		}

		res, err := h.cli().UserClient.BatchBanUser(ctx, &req)
		if err != nil {
			response.MicroErr(c, err)
			return
		}

		var message string

		if req.Type == "ban" {
			message = "禁用成功"
		} else {
			message = "启用成功"
		}

		response.MsgOK(c, res.Status, message)
	}
}

// 获取用户列表
func (h *UserHandler) List() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 链路追踪上下文
		ctx, ok := web.ContextWithSpan(c)
		if !ok {
			response.Err(c, http.StatusInternalServerError, "上下文出错")
			return
		}

		req := srv_user.GetUserListRequest{}
		if err := c.ShouldBindQuery(&req); err != nil {
			response.Err(c, http.StatusBadRequest, "参数有误")
		}

		res, err := h.cli().UserClient.GetUserList(ctx, &req)
		if err != nil {
			response.MicroErr(c, err)
			return
		}
		data := make(map[string]interface{})
		data["items"] = res.Data

		response.OK(c, data)
	}
}

// 获取管理员
func (h *UserHandler) AdminUserList() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 链路追踪上下文
		ctx, ok := web.ContextWithSpan(c)
		if !ok {
			response.Err(c, http.StatusInternalServerError, "上下文出错")
			return
		}

		req := srv_user.GetAdminUserListRequest{}
		//if err := c.ShouldBindJSON(&req); err != nil {
		//	response.Err(c, http.StatusBadRequest, "参数有误")
		//	return
		//}

		res, err := h.cli().UserClient.GetAdminUserList(ctx, &req)
		if err != nil {
			response.MicroErr(c, err)
		}

		data := make(map[string]interface{})
		data["items"] = res.Data

		response.OK(c, data)
	}
}

// 绑定管理员
func (h *UserHandler) BindAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 链路追踪上下文
		ctx, ok := web.ContextWithSpan(c)
		if !ok {
			response.Err(c, http.StatusInternalServerError, "上下文出错")
			return
		}

		req := srv_user.BindAdminRequest{}
		if err := c.ShouldBindJSON(&req); err != nil {
			response.Err(c, http.StatusBadRequest, "参数有误")
			return
		}

		res, err := h.cli().UserClient.BindAdmin(ctx, &req)
		if err != nil {
			response.MicroErr(c, err)
			return
		}

		response.MsgOK(c, res.Status, "绑定成功")
	}
}

func (h *UserHandler) ResetPassword(c *gin.Context) error {
	// 链路追踪上下文
	ctx, ok := web.ContextWithSpan(c)
	if !ok {
		return apiErr.UnknownError("上下文出错")
	}

	req := user.ResetPasswordRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		return apiErr.ParameterError("参数有误")
	}

	r := redis2.GetRedisPool().Get()
	defer r.Close()
	code, _ := redis.String(r.Do("GET", "reset-password-verify-code-"+req.Phone))
	if code == "" {
		return apiErr.UnauthorizedError("该号码还没获取或验证码已过期，请重新获取")
	}
	if code != req.VerifyCode {
		return apiErr.UnauthorizedError("验证码错误")
	}

	rpcReq := srv_user.ResetPasswordRequest{
		Phone:    req.Phone,
		Password: req.Password,
	}
	_, err := h.cli().UserClient.ResetPassword(ctx, &rpcReq)
	if err != nil {
		return apiErr.UnknownError(err.Error())
	}

	return rsp.OkMsgRsp(c, "", "重置成功")
}
