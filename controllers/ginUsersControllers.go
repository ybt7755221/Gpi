package controllers

import (
	"github.com/gin-gonic/gin"
	"gpi/dao"
	et "gpi/entities"
	"gpi/libraries/gutil"
	"gpi/service"
	"strconv"
)

type GinUsersController struct {
	serv *service.GinUsersService
}

// @Tags users表操作
// @Summary 【GetAll】根据条件获取信息
// @Description 根据条件获取信息
// @Accept html
// @Produce json
// @Param	page_num	query 	int		false	"页数，默认1"
// @Param	page_size	query 	int		false	"每夜条数，默认50"
// @Param	sort		query 	string	false	"排序。id desc, time asc"
// @Success 200 {object} SgrResp
// @Router /users [get]
func (c *GinUsersController) Find(ctx *gin.Context) {
	fieldsArr := []string{}
	//处理分页参数
	params := getPagingParams(ctx)
	//处理查询条件
	params["conditions"] = getParams(ctx, fieldsArr, et.GinUsers{})
	ginUsersList, err := c.serv.Find(params)
	if err != nil {
		resError(ctx, et.EntityFailure, err.Error())
		return
	}
	resSuccess(ctx, ginUsersList)
}

// @Tags users表操作
// @Summary 【GetOne】根据id获取信息
// @Description 根据id获取信息
// @Accept html
// @Produce json
// @Param   id		path	string 	false	"主键id"
// @Success 200 {object} SgrResp
// @Router /users/{id} [get]
func (c *GinUsersController) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	ginUsers, err := c.serv.FindById(id)
	if err != nil {
		resError(ctx, et.EntityFailure, err.Error())
	} else {
		ginUsersDao := new(dao.GinUsersDao)
		gutil.BeanUtil(ginUsersDao, ginUsers)
		resSuccess(ctx, ginUsersDao)
	}
}

// @Tags users表操作
// @Summary 【create】创建users信息
// @Description 创建users信息
// @Accept x-www-form-urlencoded
// @Produce json
// @Success 200 {object} SgrResp
// @Router /users [post]
func (c *GinUsersController) Create(ctx *gin.Context) {
	ginUsers := new(et.GinUsers)
	getPostStructData(ctx, ginUsers)
	if err := c.serv.Insert(ginUsers); err != nil {
		resError(ctx, et.EntityFailure, err.Error())
		return
	}
	resSuccess(ctx, ginUsers)
}

// @Tags users表操作
// @Summary 【update】根据id更新数据
// @Description 根据id更新数据
// @Accept x-www-form-urlencoded
// @Produce json
// @Param   id	body	string 	true	"主键更新依据此id"
// @Success 200 {object} SgrResp
// @Router /users/update-by-id [put]
func (c *GinUsersController) UpdateById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	ginUsers := new(et.GinUsers)
	getPostStructData(ctx, ginUsers)
	has, err := c.serv.UpdateById(id, ginUsers)
	if err != nil {
		resError(ctx, et.EntityFailure, err.Error())
	} else {
		if has == 0 {
			resError(ctx, et.EntityFailure, "影响行数0")
		} else {
			resSuccess(ctx, gin.H{
				"update_count": has,
			})
		}
	}
}
