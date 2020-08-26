package config

import (
	"encoding/json"
	"io/ioutil"
)

type Etcd struct {
	EtcdAddressDev []string `json:"etcdAddressDev"`
	EtcdAddress    []string `json:"etcdAddress"`
}

var G_Etcd *Etcd

func InitRegistry(fileName string) error {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	var conf Etcd

	if err := json.Unmarshal(content, &conf); err != nil {
		return err
	}

	G_Etcd = &conf
	if G_Config.Dev {
		G_Etcd.EtcdAddress = conf.EtcdAddressDev
	}

	return nil
}
