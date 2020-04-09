package controllers

import (
	"github.com/gin-gonic/gin"
	et "gpi/entities"
	"gpi/service"
	"strconv"
)

type GinContentsController struct {
	serv *service.GinContentsService
}

// @Tags contents表操作
// @Summary 【GetAll】根据条件获取信息
// @Description 根据条件获取信息
// @Accept html
// @Produce json
// @Param	page_num	query 	int		false	"页数，默认1"
// @Param	page_size	query 	int		false	"每夜条数，默认50"
// @Param	sort		query 	string	false	"排序。id desc, time asc"
// @Success 200 {object} SgrResp
// @Router /contents [get]
func (c *GinContentsController) Find(ctx *gin.Context) {
	ginContents := new(et.GinContents)
	getParamsNew(ctx, ginContents)
	pagination := new(et.Pagination)
	pagination.PageNum, _ = strconv.Atoi(ctx.Query("page_num"))
	pagination.PageSize, _ = strconv.Atoi(ctx.Query("page_size"))
	pagination.SortStr = ctx.Query("sort")
	ginContentsList, err := c.serv.Find(ginContents, pagination)
	if err != nil {
		resError(ctx, et.EntityFailure, err.Error())
		return
	}
	resSuccess(ctx, ginContentsList)
}

// @Tags contents表操作
// @Summary 【GetOne】根据id获取信息
// @Description 根据id获取信息
// @Accept html
// @Produce json
// @Param   id		path	string 	false	"主键id"
// @Success 200 {object} SgrResp
// @Router /contents/{id} [get]
func (c *GinContentsController) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	ginContents, err := c.serv.FindById(id)
	if err != nil {
		resError(ctx, et.EntityFailure, err.Error())
	} else {
		resSuccess(ctx, ginContents)
	}
}

// @Tags contents表操作
// @Summary 【create】创建contents信息
// @Description 创建contents信息
// @Accept x-www-form-urlencoded
// @Produce json
// @Success 200 {object} SgrResp
// @Router /contents [post]
func (c *GinContentsController) Create(ctx *gin.Context) {
	ginContents := new(et.GinContents)
	getPostStructData(ctx, ginContents)
	if err := c.serv.Insert(ginContents); err != nil {
		resError(ctx, et.EntityFailure, err.Error())
		return
	}
	resSuccess(ctx, ginContents)
}

// @Tags contents表操作
// @Summary 【update】根据id更新数据
// @Description 根据id更新数据
// @Accept x-www-form-urlencoded
// @Produce json
// @Param   id	body	string 	true	"主键更新依据此id"
// @Success 200 {object} SgrResp
// @Router /contents/update-by-id [put]
func (c *GinContentsController) UpdateById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	ginContents := new(et.GinContents)
	getPostStructData(ctx, ginContents)
	has, err := c.serv.UpdateById(id, ginContents)
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
