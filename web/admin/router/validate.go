package router

import (
	"log"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/simontuo/merp-go/web/admin/request/token"
	"github.com/simontuo/merp-go/web/admin/request/user"
)

func RegisterValidator() {
	//将验证方法注册到验证器中
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("SignPassword", token.SignPassword); err != nil {
			log.Fatal("登录密码验证器注册失败")
		}
		if err := v.RegisterValidation("SignVerifyCode", token.SignVerifyCode); err != nil {
			log.Fatal("登录验证码验证器注册失败")
		}
		if err := v.RegisterValidation("ConfirmPassword", user.ConfirmPassword); err != nil {
			log.Fatal("登录确认密码验证器注册失败")
		}
	}
}
