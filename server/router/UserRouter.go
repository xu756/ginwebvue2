package router

import (
	"example.com/mod/views"
	"example.com/mod/wxchat"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) *gin.Engine {
	rr := r.Group("/api/vue2/user/") // web管理路由
	rr.POST("/login", views.Login)
	rr.POST("/register1", views.Register1)
	rr.POST("/register2", views.Register2)
	rr.POST("/IsLogin", views.IsLogin)
	rr.GET("/getAccessToken", wxchat.GetAccessTokenRouter)
	return r
}
