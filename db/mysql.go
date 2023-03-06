package db

import (
	"common/config"
	commonLoger "common/loger"
	"common/utils"
	"fmt"
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"log"
	"time"
)

type Mysql struct {
	config             *config.MysqlOptions
	AutoMigrateDisable bool
	MysqlDB            *gorm.DB
}

func NewMysql(autoMigrateDisable bool, config *config.MysqlOptions) *Mysql {
	return &Mysql{config: config, AutoMigrateDisable: autoMigrateDisable}
}
func (my *Mysql) Connect() *gorm.DB {
	timeout := my.config.MysqlTimeout
	maxopenconns := my.config.MysqlMaxOpenCons
	maxidleconns := my.config.MysqlMaxIdleCons
	lifeTime := my.config.MysqlLifeTimeout
	fmt.Println(timeout, maxopenconns, maxidleconns, lifeTime)
	var link = my.config.DbUser + ":" + my.config.DbPassWord + "@tcp(" + my.config.DbHost + ":" + my.config.DbPort + ")/" + my.config.DbName + "?charset=utf8&parseTime=True&loc=Local&interpolateParams=true&timeout=" + timeout
	fmt.Println(link)
	comLoger, _ := commonLoger.NewLoger("", log.LstdFlags, func() string {
		now := time.Now()
		filename := fmt.Sprintf("my_%d%02d%02d_%02d_%02d_%02d.log",
			now.Year(),
			now.Month(),
			now.Day(),
			now.Hour(),
			now.Minute(),
			now.Second())
		return filename
	}).Init()

	newLogger := gormLogger.New(
		log.New(comLoger.Writer(), "\r\n", log.LstdFlags), // io writer)
		gormLogger.Config{
			SlowThreshold: 200000 * time.Microsecond, // 慢 SQL 阈值
			LogLevel:      gormLogger.Warn,           // Log level
			Colorful:      true,                      // 禁用彩色打印
		},
	)

	db, err := gorm.Open(gormMysql.Open(link), &gorm.Config{
		DisableAutomaticPing:                     true,
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
		PrepareStmt:                              true,
		Logger:                                   newLogger,
	})
	if err != nil {
		fmt.Println("mysql connect fail:", err.Error())
	}
	sqlDB, err := db.DB()
	sqlDB.SetConnMaxLifetime(time.Duration(lifeTime) * time.Second)
	sqlDB.SetMaxOpenConns(maxopenconns) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(maxidleconns)
	my.MysqlDB = db
	return db
}
func (my *Mysql) DB() *gorm.DB {
	t1 := utils.TimeMilliUnix()
	if my.MysqlDB == nil {
		fmt.Println("初始化mysql连接！")
		my.Connect()
	}
	sqlDB, err := my.MysqlDB.DB()
	if err != nil {
		fmt.Println("连接mysql失败，重新连接！")
		my.Connect()
	}

	err = sqlDB.Ping()
	fmt.Println("mysql status:", sqlDB.Stats())
	if err != nil {

		fmt.Println("ping失败,数据库重连")
		sqlDB.Close()
		my.Connect()
	}
	t2 := utils.TimeMilliUnix()
	fmt.Println("use time:", t2-t1)
	return my.MysqlDB
}

//AutoMigrate Mysql数据库自动同步结构体
func (my *Mysql) AutoMigrate(dst ...interface{}) {
	if my.AutoMigrateDisable {
		return
	}
	my.DB().Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(dst...)
}
