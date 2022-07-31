/**
*  @Author: xu756 (https://github.com/xu756)
*  @Email: 756334744@qq.com
*  @Date: 2022/7/28 22:55
*  @To:
 */

package upload

import (
	"crypto/sha256"
	"encoding/hex"
	"example.com/mod/cache"
	"example.com/mod/models"
	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
	"image/jpeg"
	"log"
	"os"
	"path"
	"time"
)

func Img(c *gin.Context) {
	token, _ := c.Request.Cookie("token")
	cache.RedisInit()
	useranme, _ := c.Request.Cookie("pad")
	if cache.Get(useranme.Value) != token.Value {
		c.JSON(200, gin.H{
			"type": "error",
			"msg":  "未登录",
			"is":   false,
		})
		return
	}
	// 获取上传文件
	file, err := c.FormFile("img")
	if err != nil {
		log.Panicln("无法获取上传文件")
		return
	}
	// 获取文件名
	filename := file.Filename
	// 获取文件后缀
	ext := path.Ext(filename)
	// 生成新的文件名
	h := sha256.New()
	h.Write([]byte(time.Now().String() + filename))
	filename = hex.EncodeToString(h.Sum(nil)) + ext
	// 保存文件
	err = c.SaveUploadedFile(file, "./media/upload/img/"+filename)
	if err != nil {
		log.Panicln("无法保存文件")
		return
	}
	go func() {
		file, _ := os.Open("./media/upload/img/" + filename)
		img, _ := jpeg.Decode(file)
		file.Close()
		img = resize.Resize(800, 0, img, resize.NearestNeighbor)
		out, _ := os.Create("./media/upload/img/" + filename)
		defer out.Close()
		jpeg.Encode(out, img, nil)
		log.Print("图片上传压缩成功")
	}()
	go func() {
		models.InitMysqlDB()
		var db = models.DB
		var img models.Upload
		img.Name = "img"
		img.Size = file.Size
		img.Path = "/api/get/upload/img/" + filename
		img.Type = "图片"
		db.Create(&img)
		log.Print("保存数据库成功")
	}()
	c.JSON(200, gin.H{
		"type": "success",
		"msg":  "上传成功",
		"data": gin.H{
			"url": "/api/get/upload/img/" + filename,
		},
	})

}
