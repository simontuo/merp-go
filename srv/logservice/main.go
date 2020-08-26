package main

import (
	"log"

	"github.com/micro/go-micro/v2"
	"github.com/simontuo/merp-go/config"
	"github.com/simontuo/merp-go/lib/registry"
)

const name = "go.micro.srv.log"

func main() {
	// 初始化配置
	config.Init()
	// 初始化错误ID

	service := micro.NewService(
		micro.Name(name),
		micro.Registry(registry.Etcd()),
	)

	service.Init()

	//if err := micro.RegisterSubscriber(name, service.Server(), )

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
