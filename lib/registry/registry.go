package registry

import (
	"github.com/simontuo/merp-go/config"

	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func Etcd() registry.Registry {
	//
	return etcd.NewRegistry(func(op *registry.Options) {
		op.Addrs = config.G_Etcd.EtcdAddress
	})

}
