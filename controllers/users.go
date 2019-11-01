package controllers

import (
	"github.com/gin-gonic/gin"
	et "gpi/entities"
	"gpi/models"
	"strconv"
)
//@TagName 用户模块
//@Description 用户相关接口
type Users struct {
	model *models.Users
}

// @Title Get
// @Summary 获取用户接口
// @Description 获取用户接口
// @Param	id 	query  	string  false  "id"
// @Param 	username  query string false "用户名"
// @Param	mobile	  query string false "手机号"
// @Param	offset	  query string false	"起始位置；默认0"
// @Param	limit	  query string false	"获取条数；默认20"
// @Param	sortField query string false	"排序字段；默认id"
// @param   sort	  query string false "排序顺序：1-正序，2-倒叙；默认2"
// @param   token	  query string false "验证参数"
// @Success 200 {object} ApiResonse
// @Failure 500 system err
// @router /createToken [get]
func (u *Users) Get(c *gin.Context) {
	fieldsArr := []string{"id", "username", "mobile"}
	params := getCommonParams(c)
	params["conditions"] = getParams(c, fieldsArr)
	users, err := u.model.GetUser(params)
	if err != nil {
		resError(c, 1000, err.Error())
		return
	}
	resSuccess(c, users)
}
// @Title GetId
// @Summary 根据Id获取用户接口
// @Description 根据Id获取用户接口
// @Param	id 		query  	string  true  "id"
// @Param	token	query  	string  true  "验证参数"
// @Success 200 {object} ApiResonse
// @Failure 500 system err
// @router /createToken [get, post, put, delete]
func (u *Users) GetId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := u.model.GetById(id)
	if err != nil {
		resError(c, 1000, err.Error())
		return
	}
	resSuccess(c, user)
}
// @Title Create
// @Summary 创建用户接口
// @Description 创建用户接口
// @Param	username 	query  	string  true  "用户名（拼音）"
// @Param	password 	query  	string  true  "密码"
// @Param	fullname 	query  	string  false "用户名（中文）"
// @Param	mobile	 	query  	string  true  "用户手机"
// @Param	email	 	query  	string  false "用户邮箱"
// @Param	token	 	query  	string  true  "验证参数"
// @Success 200 {object} ApiResonse
// @Failure 500 system err
// @router /createToken [post]
func (u *Users) Create(c *gin.Context) {
	userStruct := getUserBody(c)
	err := u.model.Insert(userStruct)
	if err != nil {
		resError(c, et.EntitySystemError, err.Error())
		return
	}
	resSuccess(c, userStruct)
}
// @Title Update
// @Summary 根据Id更新用户接口
// @Description 根据Id更新用户接口
// @Param	id		 	query  	string  true  "验证参数"
// @Param	username 	query  	string  false "用户名（拼音）"
// @Param	password 	query  	string  false "密码"
// @Param	fullname 	query  	string  false "用户名（中文）"
// @Param	mobile	 	query  	string  false "用户手机"
// @Param	email	 	query  	string  false "用户邮箱"
// @Param	token	 	query  	string  true  "验证参数"
// @Success 200 {object} ApiResonse
// @Failure 500 system err
// @router /createToken [put]
func (u *Users) Update(c *gin.Context) {
	userStruct := getUserBody(c)
	if c.Param("id") != c.PostForm("id") {
		resError(c, et.EntityForbidden, "Id为非法参数")
	}
	idInt, _ := strconv.Atoi(c.PostForm("id"))
	_, err := u.model.Update(idInt, userStruct)
	if err != nil {
		resError(c, et.EntitySystemError, err.Error())
		return
	}
	userStruct.Id = idInt
	resSuccess(c, userStruct)
}
/**
 * 获取user post参数
 */
func getUserBody(c *gin.Context) *et.GinUsers {
	userStruct := new(et.GinUsers)
	userStruct.Username = c.PostForm("username")
	userStruct.Password = c.PostForm("password")
	userStruct.Fullname = c.PostForm("fullname")
	userStruct.Mobile 	= c.PostForm("mobile")
	userStruct.Email 	= c.PostForm("email")
	return userStruct
}