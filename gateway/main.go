package main

import (
	"log"

	"github.com/micro/micro/v2/cmd"

	"github.com/micro/micro/v2/client/api"
	"github.com/simontuo/merp-go/config"
	"github.com/simontuo/merp-go/plugin/auth"
	"github.com/simontuo/merp-go/plugin/tracer/gateway"
)

const name = "go.micro.gateway"

func main() {
	cmd.Init()
}

func init() {
	// 初始化配置
	config.Init()

	// 鉴权验证插件
	if err := api.Register(auth.NewRegistry("./config/file/token.json")); err != nil {
		log.Fatal(err)
	}

	// 链路追踪插件
	if err := api.Register(gateway.NewRegistry()); err != nil {
		log.Fatal(err)
	}
}
