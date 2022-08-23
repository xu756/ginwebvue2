package models

import (
	"example.com/mod/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func InitMysqlDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", config.InitData.Mysql.User, config.InitData.Mysql.Password, config.InitData.Mysql.Addr, config.InitData.Mysql.Port, config.InitData.Mysql.Db)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败")
	}
	sqlDB, err := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Second)
	DB = db
	err = DB.AutoMigrate(&UserRole{}, &User{}, &Emial{}, &Upload{}, &Log{}, &ArticleCategory{}, &Article{}, &ArticleTag{}, &Menu{})
	if err != nil {
		fmt.Println("用户表创建失败")
		return
	}
}
