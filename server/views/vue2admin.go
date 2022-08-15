/**
*  @Author: xu756 (https://github.com/xu756)
*  @Email: 756334744@qq.com
*  @Date: 2022/8/8 22:31
*  @To:
 */

package views

import (
	"example.com/mod/logs"
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
	User struct {
		Id       int      `json:"id"`
		UserName string   `json:"username"` // 用户名
		Role     uint     `json:"role"`     // 角色
		Portrait string   `json:"portrait"` // 头像
		Message  []string `json:"message"`  // 消息
	}
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
	db.Where("role <= ?", role).Find(&menus)
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
	Data.User.Id = user.Id
	Data.User.UserName = user.UserName
	Data.User.Role = user.Role
	Data.User.Portrait = user.Portrait
	Data.User.Message = []string{"您有1条未读消息", "您有2条未读消息"}
	c.JSON(200, gin.H{
		"type": "success",
		"msg":  "获取数据成功",
		"data": Data,
	})
}

func GetMenu(c *gin.Context) {
	models.InitMysqlDB()
	token := c.GetHeader("token")
	var db = models.DB
	var user models.User
	var menus []models.Menu
	db.Where("token = ?", token).First(&user)
	if user.Id == 0 || user.Role != 3 {
		c.JSON(200, gin.H{
			"type": "loginout",
			"msg":  "请重新登录",
		})
		return
	}
	db.Where("role <= ?", user.Role).Find(&menus)
	// 遍历菜单，获取子菜单
	var MenuData []Menu
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
			MenuData = append(MenuData, Menu{
				Id:       menu.Id,
				Name:     menu.Name,
				Path:     menu.Path,
				Icon:     menu.Icon,
				ParentId: menu.ParentId,
				SubMenu:  sub,
			})
		}
	}
	c.JSON(200, gin.H{
		"type": "success",
		"msg":  "获取菜单成功",
		"data": MenuData,
	})
}

func UploadMenu(c *gin.Context) {
	models.InitMysqlDB()
	token := c.GetHeader("token")
	var db = models.DB
	var user models.User
	db.Where("token = ?", token).First(&user)
	if user.Id == 0 || user.Role != 3 {
		c.JSON(200, gin.H{
			"type": "loginout",
			"msg":  "请重新登录",
		})
		return
	}
	var menu SubMenu
	c.BindJSON(&menu)
	db.Where("id = ? AND path =?", menu.Id, menu.Path).Updates(models.Menu{Name: menu.Name, Icon: menu.Icon})
	c.JSON(200, gin.H{
		"type": "success",
		"msg":  "添加菜单成功",
	})
}

func GetLogs(c *gin.Context) {
	logs.Logs(logs.LogData{
		User:     "",
		Category: "",
		Ip:       c.ClientIP(),
		Type:     "普通",
		Content:  "",
	})
}
