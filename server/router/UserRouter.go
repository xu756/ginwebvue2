package router

import (
	"example.com/mod/views"
	"example.com/mod/wxchat"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) *gin.Engine {
	rr := r.Group("/api/vue2/user/")                       // web管理路由
	rr.POST("/login", views.Login)                         // 登录
	rr.POST("/register1", views.Register1)                 // 注册
	rr.POST("/register2", views.Register2)                 // 注册
	rr.POST("/IsLogin", views.IsLogin)                     // 判断是否登录
	rr.POST("/logout", views.Logout)                       // 退出登录
	rr.POST("/default", views.Default)                     // 获取默认信息，包括菜单，权限等
	rr.GET("/get/menu", views.GetMenu)                     //获取菜单数据
	rr.POST("/update/menus", views.UploadMenu)             // 更新菜单数据
	rr.GET("/getAccessToken", wxchat.GetAccessTokenRouter) // 获取access_token
	rr.GET("/getmenu", wxchat.GetMenuRouter)               // 获取公众号菜单
	rr.GET("/getlogs", views.GetLogs)                      // 获取日志
	return r

}
