/**
*  @Author: xu756 (https://github.com/xu756)
*  @Email: 756334744@qq.com
*  @Date: 2022/8/2 19:45
*  @To:
 */

package wxchat

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func Receive(c *gin.Context) {
	// 解析接收xml数据
	res := res{}
	c.BindXML(&res)
	// 判断消息类型
	log.Print(res)
	switch res.MsgType {
	case "text":
		// 文本消息
		c.XML(http.StatusOK, Text{
			ToUserName:   res.FromUserName,
			FromUserName: res.ToUserName,
			CreateTime:   time.Now().Unix(),
			MsgType:      "text",
			Content:      "您发送的是文本消息",
			XMLName:      xml.Name{},
		})
	}
}
