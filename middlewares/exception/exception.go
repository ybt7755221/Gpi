package exception

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"gpi/entities"
	"gpi/libriries/config"
	"gpi/libriries/wmail"
)

func Recover() gin.HandlerFunc{
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				go func() {
					var mailToArr []string
					conf := config.Config{}
					conf.LoadYamlConfig("errReport")
					json.Unmarshal([]byte(conf.GetString("mailto")), &mailToArr)
					msgStr := fmt.Sprintf("请求url: %s \n", c.Request.RequestURI)
					msgStr += fmt.Sprintf("请求IP: %s \n", c.ClientIP())
					msgStr += fmt.Sprintf("请求Header: %s \n", c.Request.Header)
					msgStr += fmt.Sprintf("请求时间: %s \n", time.Now().Format("2006-01-02 15:04:05"))
					msgStr += fmt.Sprintf("错误信息: %s \n", err.(string))
					wmail.SendMail(mailToArr, "系统错误", msgStr)
				}()
				c.JSON(http.StatusOK, entities.ApiResonse{http.StatusNoContent, "系统错误", gin.H{}})
				c.Abort()
			}
		}()
		c.Next()
	}
}
