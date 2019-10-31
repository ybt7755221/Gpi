package elog

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gpi/libriries/config"
	"gpi/libriries/wmail"
	"time"
)

type Elog struct {
	Eip	  string	`json:"request_ip"`
	Etime string	`json:"request_time"`
	Efile string	`json:"file_path"`
	Eline int		`json:"error_line"`
	Emsg  string	`json:"error_msg"`
}

func ErrMail(c *gin.Context, errMsg string) {
	var msgStr string
	mailToArr := config.Conf.GetStringSlice("errReport.mailto")
	subject := config.Conf.GetString("errReport.subject")
	if c == nil{
		msgStr = errMsg
	} else {
		msgStr += fmt.Sprintf("请求url: %s <br />", c.Request.RequestURI)
		msgStr += fmt.Sprintf("请求IP: %s <br />", c.ClientIP())
		msgStr += fmt.Sprintf("请求Header: %s <br />", c.Request.Header)
		msgStr += fmt.Sprintf("请求时间: %s <br />", time.Now().Format("2006-01-02 15:04:05"))
		msgStr += fmt.Sprintf("错误信息: %s <br />", errMsg)
	}
	wmail.SendMail(mailToArr, subject, msgStr)
}
