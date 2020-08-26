package main

import (
	"log"

	"github.com/micro/go-micro/v2/web"
	"github.com/simontuo/merp-go/lib/registry"
	"github.com/simontuo/merp-go/web/admin/router"

	"github.com/opentracing/opentracing-go"
	"github.com/simontuo/merp-go/config"
	"github.com/simontuo/merp-go/lib/tracer"
	webTracer "github.com/simontuo/merp-go/plugin/tracer/web"
)

const name = "go.micro.api.admin"

func main() {
	// 初始化配置
	config.Init()
	// 链路追踪
	webTracer.SetSamplingFrequency(50)
	t, io, err := tracer.NewTracer(name, "")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// 初始化路由
	r := router.Init()
	// 注册服务
	service := web.NewService(
		web.Name(name),
		web.Registry(registry.Etcd()),
	)

	config.InitWebService(service)

	service.Handle("/", r)
	// 初始化服务

	err = service.Init()

	if err != nil {
		log.Fatal(err)
	}
	// 运行服务
	if err = service.Run(); err != nil {
		log.Fatal(err)
	}
}
