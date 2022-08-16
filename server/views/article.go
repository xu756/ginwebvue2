/**
*  @Author: xu756 (https://github.com/xu756)
*  @Email: 756334744@qq.com
*  @Date: 2022/8/15 21:18
*  @To:
 */

package views

import (
	"example.com/mod/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetArticle(c *gin.Context) {

	models.InitMysqlDB()
	var db = models.DB
	var articles []models.Article
	type Data struct {
		Page     int `form:"current"`
		PageSize int `form:"size"`
	}
	var data Data
	c.BindQuery(&data)
	//db.Scopes(Paginate(data.Page, data.PageSize)).Find(&articles)
	db.Preload("Tag").Scopes(Paginate(data.Page, data.PageSize)).Find(&articles)
	var count int64
	db.Table("articles").Count(&count)
	c.JSON(200, gin.H{
		"type": "success",
		"msg":  "文章数据获取成功",
		"data": gin.H{
			"articles": articles,
			"total":    count,
		},
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

// CreateArticle 创建文章
func CreateArticle(c *gin.Context) {
	models.InitMysqlDB()
	var db = models.DB
	var article models.Article

	c.BindJSON(&article)
	var tags = article.Tag
	article.Tag = []models.ArticleTag{}
	for _, tag := range tags {
		var Tag models.ArticleTag
		db.Where(&models.ArticleTag{Name: tag.Name}).Find(&Tag)
		if Tag.Id == 0 {
			Tag.Name = tag.Name
			db.Create(&Tag)
		}
		article.Tag = append(article.Tag, Tag)

	}
	go db.Create(&article)
	c.JSON(200, gin.H{
		"type": "success",
		"msg":  "文章创建成功",
		"data": article,
	})
}
