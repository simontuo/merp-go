package main

import (
	"log"

	srv_sms "github.com/simontuo/merp-go/srv/smsservice/proto/sms"

	"github.com/micro/go-micro/v2"
	"github.com/simontuo/merp-go/lib/registry"

	"github.com/opentracing/opentracing-go"
	"github.com/simontuo/merp-go/config"
	"github.com/simontuo/merp-go/lib/tracer"
	"github.com/simontuo/merp-go/srv/smsservice/handler"

	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
)

const name = "go.micro.srv.sms"

func main() {
	// 初始化配置
	config.Init()

	// 初始化错误ID
	handler.Init(name)

	t, io, err := tracer.NewTracer(name, "")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	service := micro.NewService(
		micro.Name(name),
		micro.Registry(registry.Etcd()),
		micro.WrapHandler(ocplugin.NewHandlerWrapper(opentracing.GlobalTracer())),
	)
	if err := srv_sms.RegisterSmsServiceHandler(service.Server(), new(handler.Handler)); err != nil {
		log.Fatal(err)
	}

	service.Init()

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
