package DB

import (
	"fmt"
	"log"
	"time"

	"github.com/simontuo/merp-go/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DB() *gorm.DB {
	conf := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.G_DB.DBUser,
		config.G_DB.DBPassword,
		config.G_DB.DBHost,
		config.G_DB.DBPort,
		config.G_DB.DBDatabase,
	)

	db, err := gorm.Open("mysql", conf)
	if err != nil {
		log.Fatal(err)
	}

	// 配置连接池
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 10)
	// 开启日志
	db.LogMode(config.G_DB.DBDebug)

	return db
}
