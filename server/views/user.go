package views

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"example.com/mod/cache"
	"example.com/mod/methods"
	"example.com/mod/models"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"strconv"
	"time"
)

// 登录

func Login(c *gin.Context) {
	models.InitMysqlDB()
	var db = models.DB
	data := make(map[string]interface{}) // 定义一个map 存储客户端数据
	err := c.BindJSON(&data)             // 获取客户端数据
	if err != nil {
		log.Panicln("无法解析数据") // 如果解析失败 则报错
		return
	}
	// 对密码进行sha256加密
	h := sha256.New()
	h.Write([]byte(data["password"].(string)))
	data["password"] = hex.EncodeToString(h.Sum(nil)) // 对密码进行加密
	var user models.User
	db.Where("user_name = ?", data["username"]).First(&user)
	if user.Id == 0 {
		c.JSON(200, gin.H{
			"type": "error",
			"msg":  "用户名不存在",
		})
		return
	}
	if user.Password != data["password"] {
		c.JSON(200, gin.H{
			"type": "error",
			"msg":  "密码错误",
		})
		return
	}
	// 生成token
	h = sha512.New()
	h.Write([]byte(user.UserName + user.Password + time.Now().String()))
	token := hex.EncodeToString(h.Sum(nil))
	user.Token = token                        // 将token存入数据库
	cache.Set(user.UserName, token, 60*60*24) //redis缓存token
	user.Frequency = user.Frequency + 1
	db.Save(&user)
	c.JSON(200, gin.H{
		"type": "success",
		"msg":  "登录成功",
		"data": gin.H{
			"token":    user.Token,
			"username": user.UserName,
			"userid":   user.Id,
		},
	})
}

// Register1 注册
func Register1(c *gin.Context) {
	data := make(map[string]interface{})
	c.BindJSON(&data)
	models.InitMysqlDB()
	var db = models.DB
	var user models.User
	db.Where("user_name = ?", data["username"]).Find(&user)
	if user.Id != 0 {
		c.JSON(200, gin.H{
			"type": "error",
			"msg":  "用户名存在",
		})
		return
	}
	// 随机生成6位数字
	var code string
	for i := 0; i < 6; i++ {
		code += strconv.Itoa(rand.Intn(10)) // rand.Intn(10)生成0-9的随机数
	}
	var e = methods.EmailTo{
		To:      data["email"].(string),
		Subject: "邮箱验证",
		Body:    "您的验证码是：" + code,
	}
	if methods.SendEmail(&e) == false {
		c.JSON(200, gin.H{
			"type": "error",
			"msg":  "邮件发送失败",
		})
		return

	}
	//生产随机字符串
	str := methods.RandStr(6)
	// 将随机字符串存入redis
	cache.Set(str, code, 60) // 60秒过期
	c.JSON(200, gin.H{
		"type":    "success",
		"msg":     "验证码发送成功!",
		"randStr": str,
	})
	return
}
func Register2(c *gin.Context) {
	data := make(map[string]interface{})
	c.BindJSON(&data)
	randStr := data["randStr"].(string)
	code := data["code"].(string)
	if cache.Get(randStr) != code {
		c.JSON(200, gin.H{
			"type": "error",
			"msg":  "验证码错误",
		})
		return

	}
	//对密码进行加密
	h := sha256.New()
	h.Write([]byte(data["password"].(string)))
	data["password"] = hex.EncodeToString(h.Sum(nil)) // 对密码进行加密

	models.InitMysqlDB()
	var db = models.DB
	var user models.User
	user.UserName = data["username"].(string)
	user.Password = data["password"].(string)
	user.Role = "用户"
	user.Emial = data["email"].(string)
	//生成token
	h = sha512.New()
	h.Write([]byte(user.UserName + user.Password + time.Now().String()))
	token := hex.EncodeToString(h.Sum(nil))
	user.Token = token
	cache.Set(user.UserName, token, 60*60*24) //redis缓存token
	user.Frequency = user.Frequency + 1
	db.Create(&user)
	c.JSON(200, gin.H{
		"type": "success",
		"msg":  "注册成功",
		"data": gin.H{
			"token":    user.Token,
			"username": user.UserName,
			"userid":   user.Id,
		},
	})
}

// IsLogin 判断用户是否登录
func IsLogin(c *gin.Context) {

	token := c.GetHeader("token")
	useranme := c.GetHeader("username")
	if cache.Get(useranme) != token {
		c.JSON(200, gin.H{
			"type": "loginout",
			"msg":  "未登录",
		})
		return
	}
	c.JSON(200, gin.H{
		"type": "success",
		"msg":  "已登录",
	})
	return
}

func Logout(c *gin.Context) {
	useranme := c.GetHeader("username")
	cache.Del(useranme)
	c.JSON(200, gin.H{
		"type": "success",
		"msg":  "退出成功",
	})
	return
}
