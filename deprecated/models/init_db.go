package models

import (
	"github.com/ACking-you/byte_douyin_project/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(config.DBConnectString()), &gorm.Config{
		PrepareStmt:            true, //缓存预编译命令
		SkipDefaultTransaction: true, //禁用默认事务操作
		//Logger:                 logger.Default.LogMode(logger.Global), //打印sql语句
	})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&UserInfo{}, &Video{}, &Comment{}, &UserLogin{})
	if err != nil {
		panic(err)
	}
}
