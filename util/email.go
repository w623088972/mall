package util

import (
	"github.com/go-gomail/gomail"
	"github.com/spf13/viper"
	"log"
)

//发送电子邮件
func SendEmail(toEmail, subject, body string) error {
	fromAddress := viper.GetString("email.fromAddress")
	fromPort := viper.GetInt("email.fromPort")
	fromEmail := viper.GetString("email.fromEmail")
	fromPassword := viper.GetString("email.fromPassword")
	fromName := viper.GetString("email.fromName")

	m := gomail.NewMessage()
	m.SetAddressHeader("From", fromEmail, fromName) //发件人
	m.SetHeader("To", m.FormatAddress(toEmail, "")) //收件人
	m.SetHeader("Subject", subject)                 //主题
	m.SetBody("text/html", body)                    //内容

	d := gomail.NewPlainDialer(fromAddress, fromPort, fromEmail, fromPassword)
	if err := d.DialAndSend(m); err != nil {
		log.Println("发送失败", err)
		return err
	}
	log.Println("done.发送成功")

	return nil
}
