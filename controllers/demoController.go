package controllers

import (
	"fmt"
	et "gpi/entities"
	"gpi/libraries/mongodb"
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
	db := "system_log"
	table := "users_log"
	mongodb.InsertOne(db, table, bson.M{
		"name":     "asfasfd",
		"age":      21,
		"country":  "china",
		"password": "1231234123",
	})
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
