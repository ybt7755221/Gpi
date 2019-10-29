package controllers

import (
	"github.com/gin-gonic/gin"
	et "gpi/entities"
	"gpi/libriries/redis"
	"gpi/libriries/wmail"
	"time"
)

type Demo struct {
}

func (d *Demo) GetConf(c *gin.Context) {

}

func (d *Demo) Email(c *gin.Context) {
	err := wmail.SendMail([]string{"ybt7755221@sohu.com"}, "测试邮件", "测试呢绒")
	if err != nil {
		resError(c, et.EntityFailure, err.Error())
	}else{
		resSuccess(c, gin.H{})
	}
}

func (d *Demo) Redis(c *gin.Context) {
	redis.Cli.Set("burtyu", "30", 30*time.Second)
	res, _ := redis.Cli.Get("burtyu").Result()
	resSuccess(c, gin.H{
		"age" : res,
	})
}
