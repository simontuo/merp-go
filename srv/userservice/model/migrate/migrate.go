package main

import (
	"log"

	"github.com/simontuo/merp-go/lib/DB"
	srv_user "github.com/simontuo/merp-go/srv/userservice/proto/user"

	"github.com/simontuo/merp-go/config"
)

func main() {
	// 初始化配置
	config.Init()
	if err := config.InitDB("./config/file/DB.json", "merp-user"); err != nil {
		log.Fatal("初始化DB配置失败：" + err.Error())
	}
	db := DB.DB()

	// 迁移数据表
	if db.HasTable("users") {
		log.Print("已存在用户表")
	} else {
		db.Set("COMMENT", "用户表").CreateTable(&srv_user.User{})
		log.Print("创建用户表成功")
	}

	if db.HasTable("admin_user") {
		log.Print("已存在管理员用户表")
	} else {
		db.CreateTable(&srv_user.AdminUser{})
		log.Print("创建管理员用户表成功")
	}
}
