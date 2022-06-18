package models

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysqlDB() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败")
	}
	DB = db
	err = DB.AutoMigrate(&UserRole{}, &User{})
	if err != nil {
		fmt.Println("用户表创建失败")
		return
	}
}
