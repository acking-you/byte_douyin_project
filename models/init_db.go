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
		//Logger:                 logger.Default.LogMode(logger.Info), //打印sql语句
	})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&UserInfo{}, &Video{}, &Comment{}, &UserLogin{})
	if err != nil {
		panic(err)
	}
	//db := DB.Migrator()
	//if !db.HasTable(&UserInfo{}) {
	//	err := db.CreateTable(&UserInfo{})
	//	if err != nil {
	//		panic(err)
	//	}
	//}
	//if !db.HasTable(&Video{}) {
	//	err := db.CreateTable(&Video{})
	//	if err != nil {
	//		panic(err)
	//	}
	//}
	//if !db.HasTable(&Comment{}) {
	//	err := db.CreateTable(&Comment{})
	//	if err != nil {
	//		panic(err)
	//	}
	//}
	//if !db.HasTable(&UserLogin{}) {
	//	err := db.CreateTable(&UserLogin{})
	//	if err != nil {
	//		panic(err)
	//	}
	//}
}
