package wmail

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
	"gpi/libriries/config"
)

func SendMail(mailTo []string,subject string, body string ) error {
	conf := config.Config{}
	conf.LoadYamlConfig("email")
	m := gomail.NewMessage()
	m.SetHeader("From","gpi-system" + "<" + conf.GetString("user") + ">")
	m.SetHeader("To", mailTo...)  //发送给多个用户
	m.SetHeader("Subject", subject)  //设置邮件主题
	m.SetBody("text/html", body)     //设置邮件正文
	port, _ := conf.GetInt("port")
	d := gomail.NewDialer(
		conf.GetString("host"),
		port,
		conf.GetString("user"),
		conf.GetString("passwd"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := d.DialAndSend(m)
	return err
}