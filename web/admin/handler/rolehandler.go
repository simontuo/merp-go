package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/simontuo/merp-go/config"
	"github.com/simontuo/merp-go/helper/response"
	"github.com/simontuo/merp-go/plugin/tracer/web"
	srv_role "github.com/simontuo/merp-go/srv/permissionservice/proto/role"
)

type RoleHandler struct {
	RoleClient srv_role.RoleService
}

func (h *RoleHandler) cli() *RoleHandler {
	h.RoleClient = srv_role.NewRoleService("go.micro.srv.permission", config.WebService.Options().Service.Client())
	return h
}

// 获取角色
func (h *RoleHandler) Page() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 链路追踪上下文
		ctx, ok := web.ContextWithSpan(c)
		if !ok {
			response.Err(c, http.StatusInternalServerError, "上下文出错")
			return
		}

		req := srv_role.GetRolePageRequest{}
		if err := c.ShouldBindQuery(&req); err != nil {
			response.Err(c, http.StatusBadRequest, "参数有误")
			return
		}

		res, err := h.cli().RoleClient.GetRolePage(ctx, &req)
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

// 获取角色信息
func (h *RoleHandler) Info() gin.HandlerFunc {
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
		req := srv_role.GetRoleInfoRequest{
			Id: int32(id),
		}
		rsp, err := h.cli().RoleClient.GetRoleInfo(ctx, &req)
		if err != nil {
			response.MicroErr(c, err)
			return
		}

		response.OK(c, rsp.Data)
	}
}

// 新增角色
func (h *RoleHandler) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 链路追踪上下文
		ctx, ok := web.ContextWithSpan(c)
		if !ok {
			response.Err(c, http.StatusInternalServerError, "上下文出错")
			return
		}

		req := srv_role.StoreRoleRequest{}
		if err := c.ShouldBindJSON(&req); err != nil {
			response.Err(c, http.StatusBadRequest, "参数有误")
			return
		}
		rsp, err := h.cli().RoleClient.StoreRole(ctx, &req)
		if err != nil {
			response.MicroErr(c, err)
			return
		}

		response.MsgOK(c, rsp.Data, "新增成功")
	}
}

// 更新角色
func (h *RoleHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 链路追踪上下文
		ctx, ok := web.ContextWithSpan(c)
		if !ok {
			response.Err(c, http.StatusInternalServerError, "上下文出错")
			return
		}

		req := srv_role.UpdateRoleRequest{}
		if err := c.ShouldBindJSON(&req); err != nil {
			response.Err(c, http.StatusBadRequest, "参数有误")
			return
		}
		rsp, err := h.cli().RoleClient.UpdateRole(ctx, &req)
		if err != nil {
			response.MicroErr(c, err)
			return
		}

		response.MsgOK(c, rsp.Data, "更新成功")
	}
}

// 绑定用户
func (h *RoleHandler) BindUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 链路追踪上下文
		ctx, ok := web.ContextWithSpan(c)
		if !ok {
			response.Err(c, http.StatusInternalServerError, "上下文出错")
			return
		}

		req := srv_role.BindUserRequest{}
		if err := c.ShouldBindJSON(&req); err != nil {
			response.Err(c, http.StatusBadRequest, "参数有误")
		}
		rsp, err := h.cli().RoleClient.BindUser(ctx, &req)
		if err != nil {
			response.MicroErr(c, err)
			return
		}

		response.MsgOK(c, rsp.Status, "绑定成功")
	}
}

// 批量禁用用户
func (h *RoleHandler) BatchBan() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 链路追踪上下文
		ctx, ok := web.ContextWithSpan(c)
		if !ok {
			response.Err(c, http.StatusInternalServerError, "上下文出错")
			return
		}

		req := srv_role.BatchBanRoleRequest{}
		if err := c.ShouldBindJSON(&req); err != nil {
			response.Err(c, http.StatusBadRequest, "参数有误")
			return
		}
		res, err := h.cli().RoleClient.BatchBanRole(ctx, &req)
		if err != nil {
			response.MicroErr(c, err)
			return
		}

		var message string

		if req.Type == "ban" {
			message = "禁用采成功"
		} else {
			message = "启用企业采成功"
		}

		response.MsgOK(c, res.Status, message)
	}
}
