package db

import "gorm.io/gorm"

type AbstractDatabase interface {
	DB() *gorm.DB
	Connect() *gorm.DB
	AutoMigrate(dst ...interface{})
}
