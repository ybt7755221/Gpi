package controllers

import (
	"github.com/gin-gonic/gin"
	et "gpi/entities"
	"gpi/libraries/redis"
	"gpi/libraries/verify"
	"gpi/libraries/wmail"
	"time"
)

type DemoController struct {
}

// @Title GetToken
// @Summary 获取token接口
// @Description 获取token接口
// @Param	app_id 	query  	string  true  "验证id"
// @Param	token	query  	string  true  "验证参数"
// @Success 200 {object} ApiResonse
// @Failure 500 system err
// @router /createToken [post]
func (d *DemoController) Token(c *gin.Context) {
	var ts string
	methodStr := c.Request.Method
	if methodStr == "GET" || methodStr == "DELETE"{
		ts = c.Query("ts")
	} else {
		ts = c.PostForm("ts")
	}
	if ts != "" {
		resError(c, et.EntityParametersMissing, et.GetStatusMsg(et.EntityParametersMissing))
		return
	}
	rawStr, tokenStr := verify.GenerateToken(c)
	resSuccess(c, gin.H{
		"raw" : rawStr,
		"token" : tokenStr,
	})
}

func (d *DemoController) GetConf(c *gin.Context) {
	time.Sleep(3 * time.Second)
	resSuccess(c, c.Query("section"))
}

func (d *DemoController) Email(c *gin.Context) {
	err := wmail.SendMail([]string{"ybt7755221@sohu.com"}, "测试邮件", "测试呢绒")
	if err != nil {
		resError(c, et.EntityFailure, err.Error())
	}else{
		resSuccess(c, gin.H{})
	}
}

func (d *DemoController) Redis(c *gin.Context) {
	redis.Cache.Set("burtyu", "30", 30*time.Second)
	res, _ := redis.Cache.Get("burtyu").Result()
	resSuccess(c, gin.H{
		"age" : res,
	})
}
