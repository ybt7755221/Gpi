package controllers

import (
	"github.com/gin-gonic/gin"
	et "gpi/entities"
	"gpi/service"
	"strconv"
)
//@TagName 用户模块
//@Description 用户相关接口
type ContentsController struct {
	serv *service.ContentsService
}

func (u *ContentsController) Get(c *gin.Context) {
	fieldsArr := []string{"topic", "category", "test_time", "publish_time"}
	params := getCommonParams(c)
	params["conditions"] = getParams(c, fieldsArr)
	users, err := u.serv.Find(params)
	if err != nil {
		resError(c, et.EntitySystemError, err.Error())
		return
	}
	resSuccess(c, users)
}

func (u *ContentsController) GetId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := u.serv.FindById(id)
	if err != nil {
		resError(c, et.EntitySystemError, err.Error())
		return
	}
	resSuccess(c, user)
}

func (u *ContentsController) Create(c *gin.Context) {
	connStruct := getContentBody(c)
	err := u.serv.Insert(connStruct)
	if err != nil {
		resError(c, et.EntitySystemError, err.Error())
		return
	}
	resSuccess(c, connStruct)
}

func (u *ContentsController) Update(c *gin.Context) {
	connStruct := getContentBody(c)
	if c.Param("id") != c.PostForm("id") {
		resError(c, et.EntityForbidden, "Id为非法参数")
	}
	idInt, _ := strconv.Atoi(c.PostForm("id"))
	_, err := u.serv.UpdateById(idInt, connStruct)
	if err != nil {
		resError(c, et.EntitySystemError, err.Error())
		return
	}
	connStruct.Id = idInt
	resSuccess(c, connStruct)
}

func (u *ContentsController) Delete(c *gin.Context) {
	if c.Param("id") != c.Query("id") {
		resError(c, 1000, "Id为非法参数")
	}
	idInt, _ := strconv.Atoi(c.Query("id"))
	err := u.serv.DeleteById(idInt)
	if err != nil {
		resError(c, 1010, err.Error())
		return
	}
	resSuccess(c, gin.H{})
}
/**
 * 获取user post参数
 */
func getContentBody(c *gin.Context) *et.GinContents {
	connStruct := new(et.GinContents)
	connStruct.Topic 	  	= c.PostForm("topic")
	connStruct.Content    	= c.PostForm("content")
	connStruct.Category, _	= strconv.Atoi(c.PostForm("category"))
	connStruct.TestTime     = string2Time(c.PostForm("test_time"))
	connStruct.PulishTime 	= string2Time(c.PostForm("publish_time"))
	return connStruct
}