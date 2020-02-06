package exception

import (
	"fmt"
	"github.com/gin-gonic/gin"
	et "gpi/entities"
	"net/http"
	"time"
)

func Recover() gin.HandlerFunc{
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				go func() {
					//mailToSli := config.Conf.GetStringSlice("errReport.mailto")
					msgStr := fmt.Sprintf("请求url: %s \n", c.Request.RequestURI)
					msgStr += fmt.Sprintf("请求IP: %s \n", c.ClientIP())
					msgStr += fmt.Sprintf("请求Header: %s \n", c.Request.Header)
					msgStr += fmt.Sprintf("请求时间: %s \n", time.Now().Format("2006-01-02 15:04:05"))
					msgStr += fmt.Sprintf("错误信息: %s \n", err)
					fmt.Println(msgStr)
					//wmail.SendMail(mailToSli, config.Conf.GetString("errReport.subject"), msgStr)
				}()
				c.JSON(http.StatusOK, et.ApiResonse{et.EntitySystemError, et.GetStatusMsg(et.EntitySystemError), gin.H{}})
				c.Abort()
			}
		}()
		c.Next()
	}
}
