/**
*  @Author: xu756 (https://github.com/xu756)
*  @Email: 756334744@qq.com
*  @Date: 2022/8/1 21:17
*  @To:
 */

package wxchat

import (
	"crypto/sha1"
	"encoding/hex"
	"example.com/mod/config"
	"github.com/gin-gonic/gin"
	"sort"
)

// Verify 服务器基本配置 连接
type Verify struct {
	Signature string `form:"signature"` // 微信加密签名
	Timestamp string `form:"timestamp"` // 时间戳
	Nonce     string `form:"nonce"`     // 随机数
	Echostr   string `form:"echostr"`   // 随机字符串
}

func VerifyRouter(c *gin.Context) {
	var verify Verify
	c.BindQuery(&verify)
	// 验证
	// 拼接字符串
	str := []string{
		config.InitData.Wxcgzh.Token,
		verify.Timestamp,
		verify.Nonce,
	}
	sort.Strings(str) // 排序
	//str转换成字符串
	s := ""
	for _, v := range str {
		s += v
	}
	//sha1加密
	h := sha1.New()
	h.Write([]byte(s))
	bs := hex.EncodeToString(h.Sum([]byte("")))
	// 对比加密后的字符串
	if verify.Signature == bs {
		c.String(200, verify.Echostr)
	} else {
		c.String(200, "")
	}

}
