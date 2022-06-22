package methods

import (
	"example.com/mod/config"
	"example.com/mod/models"
	"github.com/go-gomail/gomail"
)

type EmailTo struct {
	To      string //收件人
	Subject string //主题
	Body    string //内容
}

// SendEmail 发送邮件
func SendEmail(e *EmailTo) {
	// 创建一个发送邮件的对象
	msg := gomail.NewMessage()
	// 设置发件人
	msg.SetHeader("From", msg.FormatAddress(config.Emial.From, config.Emial.Username))
	// 设置收件人
	msg.SetHeader("To", e.To)
	// 设置主题
	msg.SetHeader("Subject", e.Subject)
	// 设置内容
	msg.SetBody("text/html", e.Body)
	// 创建一个发送邮件的对象
	d := gomail.NewDialer(config.Emial.Host, config.Emial.Port, config.Emial.From, config.Emial.Password)
	// 发送邮件
	if err := d.DialAndSend(msg); err != nil {
		println("发送邮件失败", err.Error())
		return
	}
	var db = models.DB
	var emial models.Emial
	emial.To = e.To
	emial.Subject = e.Subject
	emial.Body = e.Body
	db.Create(&emial)
}
