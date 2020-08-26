package main

import (
	"log"

	"github.com/simontuo/merp-go/config"
	"github.com/simontuo/merp-go/lib/DB"
	srv_permission "github.com/simontuo/merp-go/srv/permissionservice/proto/permission"
	srv_role "github.com/simontuo/merp-go/srv/permissionservice/proto/role"
)

func main() {
	// 初始化配置
	config.Init()
	if err := config.InitDB("./config/file/DB.json", "merp-permission"); err != nil {
		log.Fatal("初始化DB配置失败：" + err.Error())
	}

	db := DB.DB()

	// 迁移数据表
	if db.HasTable("permissions") {
		log.Print("已存在权限表")
	} else {
		db.Set("COMMENT", "权限表").CreateTable(&srv_permission.Permission{})
		log.Print("创建权限表成功")
	}

	if db.HasTable("roles") {
		log.Print("已存在角色表")
	} else {
		db.Set("COMMENT", "角色表").CreateTable(&srv_role.Role{})
		log.Print("创建角色表成功")
	}

	if db.HasTable("permission_roles") {
		log.Print("已存在权限角色关系表")
	} else {
		db.Set("COMMENT", "权限角色关系表").CreateTable(&srv_permission.PermissionRole{})
		log.Print("创建权限角色关系表成功")
	}

	if db.HasTable("role_users") {
		log.Print("已存在角色用户关系表")
	} else {
		db.Set("COMMENT", "角色用户关系表").CreateTable(&srv_role.RoleUser{})
		log.Print("创建角色用户关系表成功")
	}
}
