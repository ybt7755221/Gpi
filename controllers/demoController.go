package controllers

import (
	"fmt"
	et "gpi/entities"
	"gpi/libraries/mongo"
	"gpi/libraries/redis"
	"gpi/libraries/wmail"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type DemoController struct {
}

func (d *DemoController) GetConf(c *gin.Context) {
	resSuccess(c, gin.H{})
}

func (d *DemoController) Mgo(c *gin.Context) {
	// mongo.InsertOne("system_log", "users_log", bson.M{
	// 	"name":     "saofjiasf",
	// 	"age":      31,
	// 	"country":  "china",
	// 	"password": "asfas",
	// })
	res, err := mongo.FindOne("system_log", "users_log", bson.M{"name": "saofjiasf"})
	if err != nil {
		resError(c, et.EntityFailure, err.Error())
	}
	resSuccess(c, res)
}

func (d *DemoController) Email(c *gin.Context) {
	err := wmail.SendErrMail("测试error邮件")
	if err != nil {
		resError(c, et.EntityFailure, err.Error())
	} else {
		resSuccess(c, gin.H{})
	}
}

func (d *DemoController) Redis(c *gin.Context) {
	redis.Cache.Set("burtyu", "30", 30*time.Second)
	res, err := redis.Cache.Get("burtyu").Result()
	defer redis.Cache.Close()
	fmt.Println(err)
	resSuccess(c, gin.H{
		"age": res,
	})
	return
}
