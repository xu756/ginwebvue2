/**
*  @Author: xu756 (https://github.com/xu756)
*  @Email: 756334744@qq.com
*  @Date: 2022/8/7 23:18
*  @To:	小程序基本信息获取
 */

package app

import (
	"example.com/mod/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	models.InitMysqlDB()
	var db = models.DB
	var articles []models.Article
	db.Find(&articles)
	//分页
	db.Scopes(Paginate(1, 10)).Find(&articles)
	c.JSON(200, gin.H{
		"type": "success",
		"msg":  "获取小程序基本信息成功",
	})
}

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
