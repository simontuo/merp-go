package main

import (
	"log"

	srv_tenant "github.com/simontuo/merp-go/srv/tenantservice/proto/tenant"

	"github.com/simontuo/merp-go/srv/tenantservice/handler"

	srv_department "github.com/simontuo/merp-go/srv/tenantservice/proto/department"

	"github.com/opentracing/opentracing-go"
	"github.com/simontuo/merp-go/lib/tracer"

	"github.com/micro/go-micro/v2"
	"github.com/simontuo/merp-go/config"
	"github.com/simontuo/merp-go/lib/registry"

	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
)

const name = "go.micro.srv.tenant"

func main() {
	// 初始化配置
	config.Init()

	// 初始化错误ID
	handler.Init(name)

	// 初始化数据库配置
	if err := config.InitDB("./config/file/DB.json", "merp-tenant"); err != nil {
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
	if err := srv_department.RegisterDepartmentServiceHandler(service.Server(), new(handler.DepartmentHandler)); err != nil {
		log.Fatal(err)
	}
	if err := srv_tenant.RegisterTenantServiceHandler(service.Server(), new(handler.TenantHandler)); err != nil {
		log.Fatal(err)
	}
	service.Init()

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
