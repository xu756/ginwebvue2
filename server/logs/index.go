/**
*  @Author: xu756 (https://github.com/xu756)
*  @Email: 756334744@qq.com
*  @Date: 2022/8/12 17:05
*  @To:
 */

package logs

import "example.com/mod/models"

type LogData struct {
	User     string `json:"user"`
	Role     uint   `json:"role"`
	Category string `json:"category"`
	Type     string `json:"type"`
	Content  string `json:"content"`
}

func Logs(data LogData) {
	models.InitMysqlDB()
	var db = models.DB
	var log models.Log
	var user models.User
	db.Where("user_name = ?", data.User).First(&user)
	log.User = data.User
	log.UserId = user.Id
	log.Catetory = data.Category
	log.Type = data.Type
	log.Content = data.Content
	log.Role = user.Role
	db.Create(&log)
}
