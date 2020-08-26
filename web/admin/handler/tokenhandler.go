package handler

import (
	"github.com/gomodule/redigo/redis"
	"github.com/simontuo/merp-go/helper/rsp"
	token2 "github.com/simontuo/merp-go/web/admin/request/token"

	redis2 "github.com/simontuo/merp-go/lib/redis"

	"github.com/simontuo/merp-go/config"

	"github.com/gin-gonic/gin"
	"github.com/simontuo/merp-go/helper/token"
	srv_user "github.com/simontuo/merp-go/srv/userservice/proto/user"

	apiErr "github.com/simontuo/merp-go/helper/err"
)

type TokenHandler struct {
	UserClient srv_user.UserService
}

func (h *TokenHandler) cli() *TokenHandler {
	h.UserClient = srv_user.NewUserService("go.micro.srv.user", config.WebService.Options().Service.Client())
	return h
}

func (h *TokenHandler) Sign(c *gin.Context) error {
	req := token2.SignRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		return apiErr.ParameterError("参数有误")
	}
	pass := false
	var loginUserId int32

	if req.VerifyType == "password" {
		rpcReq := srv_user.VerifyUserRequest{
			Phone:    req.Phone,
			Password: req.Password,
		}
		res, err := h.cli().UserClient.VerifyUser(c, &rpcReq)
		if err != nil {
			return apiErr.UnknownError(err.Error())
		}
		pass = res.Pass
		loginUserId = res.GetId()
	} else {
		r := redis2.GetRedisPool().Get()
		defer r.Close()
		code, _ := redis.String(r.Do("GET", "login-verify-code-"+req.Phone))
		if code == "" {
			return apiErr.UnauthorizedError("该号码还没获取或验证码已过期，请重新获取")
		}
		if code != req.VerifyCode {
			return apiErr.UnauthorizedError("验证码错误")
		} else {
			pass = true
			infoReq := srv_user.GetUserInfoRequest{Phone: req.Phone}
			user, err := h.cli().UserClient.GetUserInfo(c, &infoReq)
			if err != nil {
				return apiErr.UnknownError(err.Error())
			}
			loginUserId = user.Data.Id
		}
	}

	if pass {
		t, err := token.SignToken(req.Phone)
		if err != nil {
			return apiErr.UnknownError(err.Error())
		} else {
			redis := redis2.GetRedisPool().Get()
			defer redis.Close()

			_, err = redis.Do("SET", t, "true")
			if err != nil {
				return apiErr.UnknownError(err.Error())
			}

			data := make(map[string]interface{})
			data["userId"] = loginUserId
			data["token"] = t
			return rsp.OkRsp(c, data)
		}
	} else {
		return apiErr.UnauthorizedError("验证不通过")
	}
}

func (h *TokenHandler) Destroy(c *gin.Context) error {
	t := c.Request.Header["Token"][0]

	redis := redis2.GetRedisPool().Get()
	defer redis.Close()

	_, err := redis.Do("SET", t, "false")
	if err != nil {
		return apiErr.UnknownError(err.Error())
	}

	return rsp.OkMsgRsp(c, "", "token注销成功")
}
