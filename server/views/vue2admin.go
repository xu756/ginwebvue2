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

type Menu struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Path     string    `json:"path"`
	Icon     string    `json:"icon"`
	ParentId int       `json:"parentId"`
	SubMenu  []SubMenu `json:"subMenu"`
}

type SubMenu struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	Icon     string `json:"icon"`
	ParentId int    `json:"parentId"`
}

type DefaultData struct {
	Menu []Menu
	User models.User
}

func Default(c *gin.Context) {
	models.InitMysqlDB()
	token := c.GetHeader("token")
	var db = models.DB
	var user models.User
	var menus []models.Menu
	db.Where("token = ?", token).First(&user)
	if user.Id == 0 {
		c.JSON(200, gin.H{
			"type": "loginout",
			"msg":  "请重新登录",
		})
		return
	}
	role := user.Role
	db.Where("role = ?", role).Find(&menus)
	// 遍历菜单，获取子菜单
	var Data DefaultData
	for _, menu := range menus {
		if menu.ParentId == 0 {
			var sub []SubMenu
			var pId = menu.Id
			for _, menu := range menus {
				if menu.ParentId == pId {
					sub = append(sub, SubMenu{
						Id:       menu.Id,
						Name:     menu.Name,
						Path:     menu.Path,
						Icon:     menu.Icon,
						ParentId: menu.ParentId,
					})
				}
			}
			Data.Menu = append(Data.Menu, Menu{
				Id:       menu.Id,
				Name:     menu.Name,
				Path:     menu.Path,
				Icon:     menu.Icon,
				ParentId: menu.ParentId,
				SubMenu:  sub,
			})
		}
	}
	Data.User = user
	c.JSON(200, gin.H{
		"type": "success",
		"msg":  "获取菜单成功",
		"data": Data,
	})
}
