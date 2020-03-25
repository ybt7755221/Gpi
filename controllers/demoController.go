package controllers

import (
	et "gpi/entities"
	"gpi/libraries/httpReq"
	"gpi/libraries/mongo"
	"gpi/libraries/redis"
	"gpi/libraries/wmail"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type DemoController struct {
}

func (d *DemoController) GetConf(c *gin.Context) {
	var Ireq httpReq.IRequest
	var hreq httpReq.Request
	Ireq = hreq
	hreq.ContentType = httpReq.JSON
	hreq.Timeout = 5 * time.Second
	res, err := Ireq.Get("http://localhost:8080/createToken?ts=1573096339")
	if err != nil {
		resError(c, et.EntityFailure, err.Error())
		return
	}
	resSuccess(c, res)
}

func (d *DemoController) Mgo(c *gin.Context) {
	mongoName := "log1"
	var res interface{}
	err := mongo.FindOne("log", mongoName, bson.M{}, nil, &res)
	if err != nil {
		resError(c, et.EntityFailure, err.Error())
	}else{
		resSuccess(c, res)
	}
}

func (d *DemoController) Email(c *gin.Context) {
	err := wmail.SendErrMail("测试error邮件")
	if err != nil {
		resError(c, et.EntityFailure, err.Error())
	}else{
		resSuccess(c, gin.H{})
	}
}

func (d *DemoController) Redis(c *gin.Context) {
	redis.Cache.Set("burtyu", "30", 30*time.Second)
	res, err := redis.Cache.Get("burtyu").Result()
	defer redis.Cache.Close()
	fmt.Println(err)
	resSuccess(c, gin.H{
		"age" : res,
	})
	return
}
