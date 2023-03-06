package db

import (
	"fmt"
	"github.com/lijianjunljj/gmfcommon/config"
	"gorm.io/gorm"
)

var DB AbstractDatabase

func Init(dbType string, configFunc func() interface{}, tables ...interface{}) {
	fmt.Println("dbType", dbType)
	if dbType == "mysql" {
		conf := configFunc()
		fmt.Println("conf", conf)
		DB = NewMysql(false, conf.(*config.MysqlOptions))
		DB.AutoMigrate(tables...)
	}
}

func GetDB() *gorm.DB {
	return DB.DB()
}
