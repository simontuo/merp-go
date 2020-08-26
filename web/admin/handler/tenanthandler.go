package handler

import (
	"net/http"
	"strconv"

	"github.com/simontuo/merp-go/helper/helper"

	"github.com/gin-gonic/gin"
	"github.com/simontuo/merp-go/config"
	"github.com/simontuo/merp-go/helper/response"
	"github.com/simontuo/merp-go/plugin/tracer/web"
	srv_tenant "github.com/simontuo/merp-go/srv/tenantservice/proto/tenant"
)

type TenantHandler struct {
	TenantClient srv_tenant.TenantService
}

func (h *TenantHandler) cli() *TenantHandler {
	h.TenantClient = srv_tenant.NewTenantService("go.micro.srv.tenant", config.WebService.Options().Service.Client())
	return h
}

// 获取租户
func (h *TenantHandler) Page() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 链路追踪上下文
		ctx, ok := web.ContextWithSpan(c)
		if !ok {
			response.Err(c, http.StatusInternalServerError, "上下文出错")
			return
		}

		req := srv_tenant.GetTenantPageRequest{}
		if err := c.ShouldBindQuery(&req); err != nil {
			response.Err(c, http.StatusBadRequest, "参数有误")
			return
		}

		res, err := h.cli().TenantClient.GetTenantPage(ctx, &req)
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

// 新增租户
func (h *TenantHandler) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 链路追踪上下文
		ctx, ok := web.ContextWithSpan(c)
		if !ok {
			response.Err(c, http.StatusInternalServerError, "上下文出错")
			return
		}

		req := srv_tenant.StoreTenantRequest{}
		if err := c.ShouldBindJSON(&req.Form); err != nil {
			response.Err(c, http.StatusBadRequest, "参数有误")
			return
		}

		res, err := h.cli().TenantClient.StoreTenant(ctx, &req)
		if err != nil {
			response.MicroErr(c, err)
			return
		}

		response.MsgOK(c, res.Data, "新增成功")
	}
}

// 新增租户
func (h *TenantHandler) Info() gin.HandlerFunc {
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
		req := srv_tenant.GetTenantInfoRequest{
			Id: int32(id),
		}
		rsp, err := h.cli().TenantClient.GetTenantInfo(ctx, &req)
		if err != nil {
			response.MicroErr(c, err)
			return
		}
		data := helper.Translate(rsp.Data)

		response.OK(c, data)
	}
}

// 更新租户
func (h *TenantHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 链路追踪上下文
		ctx, ok := web.ContextWithSpan(c)
		if !ok {
			response.Err(c, http.StatusInternalServerError, "上下文出错")
			return
		}

		req := srv_tenant.UpdateTenantRequest{}
		if err := c.ShouldBindJSON(&req.Form); err != nil {
			response.Err(c, http.StatusBadRequest, "参数有误")
		}
		rsp, err := h.cli().TenantClient.UpdateTenant(ctx, &req)
		if err != nil {
			response.MicroErr(c, err)
			return
		}

		response.MsgOK(c, rsp.Data, "更新成功")
	}
}
