package main

import (
	"log"

	srv_tenant "github.com/simontuo/merp-go/srv/tenantservice/proto/tenant"
	srv_tenantuser "github.com/simontuo/merp-go/srv/tenantservice/proto/tenantuser"

	srv_department "github.com/simontuo/merp-go/srv/tenantservice/proto/department"

	"github.com/simontuo/merp-go/config"

	"github.com/simontuo/merp-go/lib/DB"
)

func main() {
	// 初始化配置
	config.Init()
	if err := config.InitDB("./config/file/DB.json", "merp-tenant"); err != nil {
		log.Fatal("初始化DB配置失败：" + err.Error())
	}

	db := DB.DB()

	// 迁移数据表
	if db.HasTable("departments") {
		log.Print("已存在部门表")
	} else {
		db.Set("COMMENT", "部门表").CreateTable(&srv_department.Department{})
		log.Print("创建部门表成功")
	}

	if db.HasTable("tenants") {
		log.Print("已存在租户表")
	} else {
		db.Set("COMMENT", "租户表").CreateTable(&srv_tenant.Tenant{})
		log.Print("创建租户表成功")
	}

	if db.HasTable("tenant_users") {
		log.Print("已存在租户用户关系表")
	} else {
		db.Set("COMMENT", "租户用户关系表").CreateTable(&srv_tenantuser.TenantUser{})
		log.Print("创建租户用户关系表成功")
	}
}
