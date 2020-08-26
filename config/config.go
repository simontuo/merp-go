package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/go-acme/lego/v3/log"
)

type Config struct {
	Dev bool `json:"dev"`
}

var G_Config *Config

func Init() {
	conf := Config{
		Dev: true,
	}
	G_Config = &conf

	// 初始化服务注册配置
	if err := InitRegistry("./config/file/registry.json"); err != nil {
		log.Fatal("初始化服务注册配置失败：" + err.Error())
	}
	log.Print("初始化服务注册配置成功")

	if err := InitRedis("./config/file/redis.json"); err != nil {
		log.Fatal("初始化redis配置失败：" + err.Error())
	}
	log.Print("初始化redis配置成功")

	if err := InitSMS("./config/file/sms.json"); err != nil {
		log.Fatal("初始化sms配置失败：" + err.Error())
	}
	log.Print("初始化sms配置成功")

	// 初始化token配置
	if err := InitToken("./config/file/token.json"); err != nil {
		log.Fatal("初始化token配置失败：" + err.Error())
	}
	log.Print("初始化token配置成功")
}

type Redis struct {
	Host        string `json:"host"`
	Port        string `json:"port"`
	Password    string `json:"password"`
	DevHost     string `json:"devHost"`
	DevPort     string `json:"devPort"`
	DevPassword string `json:"devPassword"`
}

var G_Redis *Redis

func InitRedis(fileName string) error {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	var conf Redis

	if err := json.Unmarshal(content, &conf); err != nil {
		return err
	}

	G_Redis = &conf
	if G_Config.Dev {
		G_Redis.Host = conf.DevHost
		G_Redis.Password = conf.DevPassword
		G_Redis.Port = conf.DevPort
	}

	return nil
}
