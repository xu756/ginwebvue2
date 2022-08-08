/**
*  @Author: xu756 (https://github.com/xu756)
*  @Email: 756334744@qq.com
*  @Date: 2022/8/7 23:19
*  @To:	小程序路由
 */

package router

import (
	"example.com/mod/app"
	"github.com/gin-gonic/gin"
)

func App(r *gin.Engine) *gin.Engine {
	rr := r.Group("/api/app/") // web管理路由
	rr.GET("/default", app.Index)
	return r
}
