package main

import (
	"log"

	"github.com/opentracing/opentracing-go"
	"github.com/simontuo/merp-go/lib/tracer"

	"github.com/simontuo/merp-go/srv/userservice/handler"
	srv_user "github.com/simontuo/merp-go/srv/userservice/proto/user"

	"github.com/micro/go-micro/v2"
	"github.com/simontuo/merp-go/config"
	"github.com/simontuo/merp-go/lib/registry"

	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
)

const name = "go.micro.srv.user"

func main() {
	// 初始化配置
	config.Init()

	// 初始化错误ID
	handler.Init(name)

	// 初始化数据库配置
	if err := config.InitDB("./config/file/DB.json", "merp-user"); err != nil {
		log.Fatal("初始化DB配置失败：" + err.Error())
	}
	log.Print("初始化DB配置成功")

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
	if err := srv_user.RegisterUserServiceHandler(service.Server(), new(handler.Handler)); err != nil {
		log.Fatal(err)
	}

	service.Init()

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
