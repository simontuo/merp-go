package auth

import (
	"log"
	"net/http"

	"github.com/gomodule/redigo/redis"
	redis2 "github.com/simontuo/merp-go/lib/redis"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2/errors"
	"github.com/micro/micro/v2/plugin"
	"github.com/simontuo/merp-go/config"
	token2 "github.com/simontuo/merp-go/helper/token"
	"github.com/simontuo/merp-go/plugin/util/response"
)

const id = "micro.go.plugin.auth"

type Auth struct {
	options  Options
	disable  bool
	passPath []string
}

func NewRegistry(file string) plugin.Plugin {
	// 初始化token配置
	if err := config.InitToken(file); err != nil {
		log.Fatal("初始化token配置失败：" + err.Error())
	}
	a := Auth{
		passPath: config.G_Token.PassPath,
	}
	a.options.responseHandler = response.DefaultResponseHandler

	return plugin.NewPlugin(
		plugin.WithName("Auth"),
		plugin.WithFlag(&cli.BoolFlag{
			Name:  "auth_disable",
			Usage: "disable auth",
			Value: false,
		}),
		plugin.WithInit(func(ctx *cli.Context) error {
			a.disable = ctx.Bool("auth_disable")

			return nil
		}),
		plugin.WithHandler(a.handler),
	)
}

func (a *Auth) handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if a.disable {
			h.ServeHTTP(w, r)
			return
		}
		pass := false
		path := r.URL.Path
		for _, v := range a.passPath {
			if path == v {
				pass = true
				break
			}
		}
		if pass {
			h.ServeHTTP(w, r)
			return
		}

		// 验证token
		if token, ok := r.Header["Token"]; ok {
			redisClient := redis2.GetRedisPool().Get()
			defer redisClient.Close()
			active, _ := redis.String(redisClient.Do("GET", token[0]))
			if active != "true" {
				a.options.responseHandler(w, r, errors.Unauthorized(id, "token已注销，请重新获取"))
				return
			}

			if claims, err := token2.ParseToken(token[0]); err != nil {
				a.options.responseHandler(w, r, errors.Unauthorized(id, err.Error()))
				return
			} else {
				r.Header.Set("loginUser", claims.UserPhone)
				h.ServeHTTP(w, r)
			}
		} else {
			a.options.responseHandler(w, r, errors.Unauthorized(id, "不存在token"))
			return
		}
	})
}
