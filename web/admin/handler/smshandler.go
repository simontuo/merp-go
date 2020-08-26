package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/simontuo/merp-go/config"
	apiErr "github.com/simontuo/merp-go/helper/err"
	"github.com/simontuo/merp-go/helper/rsp"
	"github.com/simontuo/merp-go/plugin/tracer/web"
	srv_sms "github.com/simontuo/merp-go/srv/smsservice/proto/sms"
	"github.com/simontuo/merp-go/web/admin/request/sms"
)

type SmsHandler struct {
	SmsClient srv_sms.SmsService
}

func (h *SmsHandler) cli() *SmsHandler {
	h.SmsClient = srv_sms.NewSmsService("go.micro.srv.sms", config.WebService.Options().Service.Client())
	return h
}

func (h *SmsHandler) LoginVerifyCode(c *gin.Context) error {
	// 链路追踪上下文
	ctx, ok := web.ContextWithSpan(c)
	if !ok {
		return apiErr.UnknownError("上下文出错")
	}
	req := sms.LoginVerifyCodeRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		return apiErr.ParameterError("参数有误")
	}
	rpcReq := srv_sms.GetLoginVerifyCodeRequest{
		Phone: req.Phone,
	}

	_, err := h.cli().SmsClient.GetLoginVerifyCode(ctx, &rpcReq)
	if err != nil {
		return apiErr.UnknownError(err.Error())
	}

	return rsp.OkMsgRsp(c, "", "发送成功")
}

func (h *SmsHandler) ForgetPasswordVerifyCode(c *gin.Context) error {
	// 链路追踪上下文
	ctx, ok := web.ContextWithSpan(c)
	if !ok {
		return apiErr.UnknownError("上下文出错")
	}
	req := sms.LoginVerifyCodeRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		return apiErr.ParameterError("参数有误")
	}
	rpcReq := srv_sms.GetForgetPasswordCodeRequest{
		Phone: req.Phone,
	}

	_, err := h.cli().SmsClient.GetForgetPasswordCode(ctx, &rpcReq)
	if err != nil {
		return apiErr.UnknownError(err.Error())
	}

	return rsp.OkMsgRsp(c, "", "发送成功")
}
