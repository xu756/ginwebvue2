package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default() // 初始化路由
	//404
	r.NoRoute(func(c *gin.Context) {
		// 获取请求的url
		url := c.Request.URL.String()
		c.JSON(404, gin.H{
			"code": 404,
			"msg":  "404 not found",
			"url":  url,
		},
		)
	})
	UserRouter(r)                              // 用户模块
	r.SetTrustedProxies([]string{"127.0.0.1"}) //设置代理
	err := r.Run(":7000")                      // 监听端口
	fmt.Println("启动成功")
	if err != nil {
		fmt.Println("启动失败") // 启动失败
		return
	}
}
