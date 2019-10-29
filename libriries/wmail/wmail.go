package wmail

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
	"gpi/libriries/config"
)

func SendMail(mailTo []string,subject string, body string ) error {
	conf := config.Conf.GetStringMap("email")
	m := gomail.NewMessage()
	m.SetHeader("From","gpi-system" + "<" + conf["user"].(string) + ">")
	m.SetHeader("To", mailTo...)  //发送给多个用户
	m.SetHeader("Subject", subject)  //设置邮件主题
	m.SetBody("text/html", body)     //设置邮件正文
	d := gomail.NewDialer(
		conf["host"].(string),
		conf["port"].(int),
		conf["user"].(string),
		conf["passwd"].(string))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := d.DialAndSend(m)
	return err
}