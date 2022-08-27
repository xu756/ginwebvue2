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
)

func Index(c *gin.Context) {
	models.InitMysqlDB()
	var db = models.DB
	var articles []models.Article
	db.Where("category = '深度校园'").Order("-id").Limit(10).Find(&articles)
	c.JSON(200, gin.H{
		"type": "success",
		"msg":  "获取小程序基本信息成功",
		"data": articles,
	})
}
