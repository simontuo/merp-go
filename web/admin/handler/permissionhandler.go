package handler

import (
	"net/http"
	"strconv"

	"github.com/simontuo/merp-go/helper/helper"

	"github.com/simontuo/merp-go/helper/tree"

	"github.com/gin-gonic/gin"
	"github.com/simontuo/merp-go/config"
	"github.com/simontuo/merp-go/helper/response"
	"github.com/simontuo/merp-go/plugin/tracer/web"
	srv_permission "github.com/simontuo/merp-go/srv/permissionservice/proto/permission"
)

type PermissionHandler struct {
	PermissionClient srv_permission.PermissionService
}

func (h *PermissionHandler) cli() *PermissionHandler {
	h.PermissionClient = srv_permission.NewPermissionService("go.micro.srv.permission", config.WebService.Options().Service.Client())
	return h
}

// 获取权限
func (h *PermissionHandler) Page() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 链路追踪上下文
		ctx, ok := web.ContextWithSpan(c)
		if !ok {
			response.Err(c, http.StatusInternalServerError, "上下文出错")
			return
		}

		req := srv_permission.GetPermissionPageRequest{}
		if err := c.ShouldBindQuery(&req); err != nil {
			response.Err(c, http.StatusBadRequest, "参数有误")
			return
		}

		res, err := h.cli().PermissionClient.GetPermissionPage(ctx, &req)
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

// 获取权限
func (h *PermissionHandler) Info() gin.HandlerFunc {
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
		req := srv_permission.GetPermissionInfoRequest{
			Id: int32(id),
		}
		res, err := h.cli().PermissionClient.GetPermissionInfo(ctx, &req)
		if err != nil {
			response.MicroErr(c, err)
			return
		}
		data := helper.Translate(res.Data)
		response.OK(c, data)
	}
}

// 获取权限树
func (h *PermissionHandler) Tree() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 链路追踪上下文
		ctx, ok := web.ContextWithSpan(c)
		if !ok {
			response.Err(c, http.StatusInternalServerError, "上下文出错")
			return
		}

		req := srv_permission.GetPermissionTreeRequest{}

		res, err := h.cli().PermissionClient.GetPermissionTree(ctx, &req)
		if err != nil {
			response.MicroErr(c, err)
			return
		}

		data := res.Data
		if len(data) < 1 {
			data = []*srv_permission.Permission{}
		}

		t := tree.BuildPermissionTree(data)

		response.OK(c, tree.BuildPTree(-1, t))
	}
}

// 新增权限
func (h *PermissionHandler) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 链路追踪上下文
		ctx, ok := web.ContextWithSpan(c)
		if !ok {
			response.Err(c, http.StatusInternalServerError, "上下文出错")
			return
		}

		req := srv_permission.StorePermissionRequest{}
		if err := c.ShouldBindJSON(&req.Form); err != nil {
			response.Err(c, http.StatusBadRequest, "参数有误")
		}

		res, err := h.cli().PermissionClient.StorePermission(ctx, &req)
		if err != nil {
			response.MicroErr(c, err)
			return
		}

		response.MsgOK(c, res.Data, "新增成功")
	}
}

// 更新权限
func (h *PermissionHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 链路追踪上下文
		ctx, ok := web.ContextWithSpan(c)
		if !ok {
			response.Err(c, http.StatusInternalServerError, "上下文出错")
			return
		}

		req := srv_permission.UpdatePermissionRequest{}
		if err := c.ShouldBindJSON(&req.Form); err != nil {
			response.Err(c, http.StatusBadRequest, "参数有误")
		}

		res, err := h.cli().PermissionClient.UpdatePermission(ctx, &req)
		if err != nil {
			response.MicroErr(c, err)
			return
		}

		response.MsgOK(c, res.Data, "更新成功")
	}
}

// 批量禁用权限
func (h *PermissionHandler) BatchBan() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 链路追踪上下文
		ctx, ok := web.ContextWithSpan(c)
		if !ok {
			response.Err(c, http.StatusInternalServerError, "上下文出错")
			return
		}

		var req srv_permission.BatchBanPermissionRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			response.Err(c, http.StatusBadRequest, "参数有误")
			return
		}

		res, err := h.cli().PermissionClient.BatchBanPermission(ctx, &req)
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
