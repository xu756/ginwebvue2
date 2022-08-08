/**
*  @Author: xu756 (https://github.com/xu756)
*  @Email: 756334744@qq.com
*  @Date: 2022/8/8 22:31
*  @To:
 */

package views

import (
	"example.com/mod/models"
	"github.com/gin-gonic/gin"
)

func Getmenu(c *gin.Context) {
	token := c.GetHeader("token")
	var db = models.DB
	var users models.User
	var menus []models.Menu
	db.Where("token = ?", token).First(&users)
	if users.Id == 0 {
		c.JSON(200, gin.H{
			"type": "loginout",
			"msg":  "请重新登录",
		})
		return
	}
	role := users.Role
	db.Where("role = ?", role).Find(&menus)
	c.JSON(200, gin.H{
		"type": "success",
		"msg":  "获取菜单成功",
		"data": menus,
	})
}
