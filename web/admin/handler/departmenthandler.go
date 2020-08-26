package handler

import (
	"net/http"
	"strconv"

	"github.com/simontuo/merp-go/helper/tree"

	srv_tenant "github.com/simontuo/merp-go/srv/tenantservice/proto/department"

	"github.com/gin-gonic/gin"
	"github.com/simontuo/merp-go/config"
	"github.com/simontuo/merp-go/helper/response"
	"github.com/simontuo/merp-go/plugin/tracer/web"
)

type DepartmentHandler struct {
	DepartmentClient srv_tenant.DepartmentService
}

func (h *DepartmentHandler) cli() *DepartmentHandler {
	h.DepartmentClient = srv_tenant.NewDepartmentService("go.micro.srv.tenant", config.WebService.Options().Service.Client())
	return h
}

// 获取部门树
func (h *DepartmentHandler) Tree() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 链路追踪上下文
		ctx, ok := web.ContextWithSpan(c)
		if !ok {
			response.Err(c, http.StatusInternalServerError, "上下文出错")
			return
		}

		req := srv_tenant.GetDepartmentTreeRequest{}
		res, err := h.cli().DepartmentClient.GetDepartmentTree(ctx, &req)
		if err != nil {
			response.MicroErr(c, err)
			return
		}

		data := res.Data
		if len(data) < 1 {
			data = []*srv_tenant.Department{}
		}

		t := tree.BuildDepartmentTree(data)

		response.OK(c, tree.BuildTree(-1, t))
	}
}

// 新增部门
func (h *DepartmentHandler) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 链路追踪上下文
		ctx, ok := web.ContextWithSpan(c)
		if !ok {
			response.Err(c, http.StatusInternalServerError, "上下文出错")
			return
		}

		req := srv_tenant.StoreDepartmentRequest{}
		if err := c.ShouldBindJSON(&req.Form); err != nil {
			response.Err(c, http.StatusBadRequest, "参数有误")
			return
		}

		res, err := h.cli().DepartmentClient.StoreDepartment(ctx, &req)
		if err != nil {
			response.MicroErr(c, err)
		}

		response.MsgOK(c, res.Data, "新增成功")
	}
}

// 部门信息
func (h *DepartmentHandler) Info() gin.HandlerFunc {
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
		req := srv_tenant.GetDepartmentInfoRequest{
			Id: int32(id),
		}

		res, err := h.cli().DepartmentClient.GetDepartmentInfo(ctx, &req)

		if err != nil {
			response.MicroErr(c, err)
		}

		response.OK(c, res.Data)
	}
}

// 部门信息
func (h *DepartmentHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 链路追踪上下文
		ctx, ok := web.ContextWithSpan(c)
		if !ok {
			response.Err(c, http.StatusInternalServerError, "上下文出错")
			return
		}

		req := srv_tenant.UpdateDepartmentRequest{}
		if err := c.ShouldBindJSON(&req.Form); err != nil {
			response.Err(c, http.StatusBadRequest, "参数有误")
		}

		res, err := h.cli().DepartmentClient.UpdateDepartment(ctx, &req)
		if err != nil {
			response.MicroErr(c, err)
		}

		response.MsgOK(c, res.Data, "更新成功")
	}
}

// 部门禁用
func (h *DepartmentHandler) BatchBan() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 链路追踪上下文
		ctx, ok := web.ContextWithSpan(c)
		if !ok {
			response.Err(c, http.StatusInternalServerError, "上下文出错")
			return
		}

		req := srv_tenant.BatchBanDepartmentRequest{}
		if err := c.ShouldBindJSON(&req); err != nil {
			response.Err(c, http.StatusBadRequest, "参数有误")
		}

		res, err := h.cli().DepartmentClient.BatchBanDepartment(ctx, &req)
		if err != nil {
			response.MicroErr(c, err)
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
