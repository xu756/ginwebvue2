/**
*  @Author: xu756 (https://github.com/xu756)
*  @Email: 756334744@qq.com
*  @Date: 2022/8/1 20:33
*  @To:
 */

package router

import (
	"example.com/mod/wxchat"
	"github.com/gin-gonic/gin"
)

func Wx(r *gin.Engine) *gin.Engine {
	rr := r.Group("/api/wx/") // web管理路由
	rr.GET("/token", wxchat.VerifyRouter)
	rr.POST("/token", wxchat.Receive)
	rr.GET("/getAccessToken", wxchat.GetAccessTokenRouter)
	return r
}
