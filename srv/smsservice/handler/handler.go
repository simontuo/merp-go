package handler

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	"github.com/simontuo/merp-go/config"

	"github.com/micro/go-micro/errors"
	redis2 "github.com/simontuo/merp-go/lib/redis"
	sms2 "github.com/simontuo/merp-go/lib/sms"
	srv_sms "github.com/simontuo/merp-go/srv/smsservice/proto/sms"
)

type Handler struct {
}

var ErrId string

func Init(name string) {
	ErrId = name
}

func getRandCode() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(1000000)
}

func (h *Handler) GetForgetPasswordCode(ctx context.Context, req *srv_sms.GetForgetPasswordCodeRequest, rsp *srv_sms.GetForgetPasswordCodeResponse) error {
	code := getRandCode()

	sms := sms2.HuaWeyYunSms{}
	sms.Body.To = req.Phone
	sms.Body.TemplateParas = []string{strconv.Itoa(code)}
	sms.Body.TemplateId = config.G_SMS.ResetPasswordTemplateId

	res, err := sms.Send()
	if err != nil {
		return errors.InternalServerError(ErrId, "短信发送失败")
	}
	if res.Code != "000000" {
		return errors.InternalServerError(ErrId, res.Description)
	}

	// redis 缓存reset
	redis := redis2.GetRedisPool().Get()
	defer redis.Close()
	_, err = redis.Do("SET", "reset-password-verify-code-"+req.Phone, code, "EX", "300")
	if err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}

	rsp.Status = "success"

	return nil
}

func (h *Handler) GetLoginVerifyCode(ctx context.Context, req *srv_sms.GetLoginVerifyCodeRequest, rsp *srv_sms.GetLoginVerifyCodeResponse) error {
	code := getRandCode()

	sms := sms2.HuaWeyYunSms{}
	sms.Body.To = req.Phone
	sms.Body.TemplateParas = []string{strconv.Itoa(code)}
	sms.Body.TemplateId = config.G_SMS.LoginTemplateId

	res, err := sms.Send()
	if err != nil {
		return errors.InternalServerError(ErrId, "短信发送失败")
	}
	if res.Code != "000000" {
		return errors.InternalServerError(ErrId, res.Description)
	}

	// redis 缓存
	redis := redis2.GetRedisPool().Get()
	defer redis.Close()
	_, err = redis.Do("SET", "login-verify-code-"+req.Phone, code, "EX", "300")
	if err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}

	rsp.Status = "success"

	return nil
}
