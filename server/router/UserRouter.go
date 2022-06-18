package router

import "github.com/gin-gonic/gin"

func UserRouter(r *gin.Engine) *gin.Engine {
	rr := r.Group("/api/web/user/") // web管理路由
	rr.GET("/", func(c *gin.Context) {
		c.String(200, "user")

	})

	return r
}
