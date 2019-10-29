package elog

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"gpi/libriries/config"
	"gpi/libriries/wmail"
)

type Elog struct {
	Eip	  string	`json:"request_ip"`
	Etime string	`json:"request_time"`
	Efile string	`json:"file_path"`
	Eline int		`json:"error_line"`
	Emsg  string	`json:"error_msg"`
}

func ErrMail(c *gin.Context, errMsg string) {
	var mailToArr []string
	configErr := config.Config{}
	configErr.LoadYamlConfig("errReport")
	json.Unmarshal([]byte(configErr.GetString("mailto")), &mailToArr)
	msgStr := fmt.Sprintf("请求url: %s \n", c.Request.RequestURI)
	msgStr += fmt.Sprintf("请求IP: %s \n", c.ClientIP())
	msgStr += fmt.Sprintf("请求Header: %s \n", c.Request.Header)
	msgStr += fmt.Sprintf("请求时间: %s \n", time.Now().Format("2006-01-02 15:04:05"))
	msgStr += fmt.Sprintf("错误信息: %s \n", errMsg)
	wmail.SendMail(mailToArr, "系统错误", msgStr)
}
