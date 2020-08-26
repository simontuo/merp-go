package main

import (
	"log"

	"github.com/simontuo/merp-go/srv/permissionservice/handler"
	srv_permission "github.com/simontuo/merp-go/srv/permissionservice/proto/permission"
	srv_role "github.com/simontuo/merp-go/srv/permissionservice/proto/role"

	"github.com/opentracing/opentracing-go"
	"github.com/simontuo/merp-go/lib/tracer"

	"github.com/micro/go-micro/v2"
	"github.com/simontuo/merp-go/config"
	"github.com/simontuo/merp-go/lib/registry"

	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
)

const name = "go.micro.srv.permission"

func main() {
	// 初始化配置
	config.Init()

	// 初始化错误ID
	handler.Init(name)

	// 初始化数据库配置
	if err := config.InitDB("./config/file/DB.json", "merp-permission"); err != nil {
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
	if err := srv_permission.RegisterPermissionServiceHandler(service.Server(), new(handler.PermissionHandler)); err != nil {
		log.Fatal(err)
	}
	if err := srv_role.RegisterRoleServiceHandler(service.Server(), new(handler.RoleHandler)); err != nil {
		log.Fatal(err)
	}
	service.Init()

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
