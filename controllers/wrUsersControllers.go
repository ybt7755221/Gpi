package controllers

import (
	et "gpi/entities"
	"gpi/service"
	"github.com/gin-gonic/gin"
	"strconv"
)
type WrUsersController struct {
	serv *service.WrUsersService
}
// @Tags wrUsers表操作
// @Summary 【GetAll】根据条件获取信息
// @Description 根据条件获取信息
// @Accept html
// @Produce json
// @Param	page_num	query 	int		false	"页数，默认1"
// @Param	page_size	query 	int		false	"每夜条数，默认50"
// @Param	sort		query 	string	false	"排序。id desc, time asc"
// @Success 200 {object} SgrResp
// @Router /wrUsers [get]
func (c *WrUsersController) Find(ctx *gin.Context) {
	wrUsers := new(et.WrUsers)
	getParamsNew(ctx, wrUsers)
	pagination := new(et.Pagination)
	pagination.PageNum, _ = strconv.Atoi(ctx.Query("page_num"))
	pagination.PageSize, _ = strconv.Atoi(ctx.Query("page_size"))
	pagination.SortStr = ctx.Query("sort")
	wrUsersList, err := c.serv.Find(wrUsers, pagination)
	if err != nil {
		resError(ctx, et.EntityFailure, err.Error())
		return
	}
	resSuccess(ctx, wrUsersList)
}

// @Tags wrUsers表操作
// @Summary 【GetAll】根据条件获取信息
// @Description 根据条件获取信息
// @Accept html
// @Produce json
// @Param	page_num	query 	int		false	"页数，默认1"
// @Param	page_size	query 	int		false	"每夜条数，默认50"
// @Param	sort		query 	string	false	"排序。id desc, time asc"
// @Success 200 {object} SgrResp
// @Router /wrUsers [get]
func (c *WrUsersController) FindPaging(ctx *gin.Context) {
	wrUsers := new(et.WrUsers)
	getParamsNew(ctx, wrUsers)
	pagination := new(et.Pagination)
	pagination.PageNum, _ = strconv.Atoi(ctx.Query("page_num"))
	pagination.PageSize, _ = strconv.Atoi(ctx.Query("page_size"))
	pagination.SortStr = ctx.Query("sort")
	wrUsersList, err := c.serv.FindPaging(wrUsers, pagination)
	if err != nil {
		resError(ctx, et.EntityFailure, err.Error())
		return
	}
	resSuccess(ctx, wrUsersList)
}
// @Tags wrUsers表操作
// @Summary 【GetOne】根据id获取信息
// @Description 根据id获取信息
// @Accept html
// @Produce json
// @Param   id		path	string 	false	"主键id"
// @Success 200 {object} SgrResp
// @Router /wrUsers/{id} [get]
func (c *WrUsersController) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	wrUsers, err := c.serv.FindById(id)
	if err != nil {
		resError(ctx, et.EntityFailure, err.Error())
	}else{
		resSuccess(ctx, wrUsers)
	}
}
// @Tags wrUsers表操作
// @Summary 【create】创建wrUsers信息
// @Description 创建wrUsers信息
// @Accept x-www-form-urlencoded
// @Produce json
// @Success 200 {object} SgrResp
// @Router /wrUsers [post]
func (c *WrUsersController) Create(ctx *gin.Context) {
	wrUsers := new(et.WrUsers)
	getPostStructData(ctx, wrUsers)
	if err := c.serv.Insert(wrUsers); err != nil {
		resError(ctx, et.EntityFailure, err.Error())
		return
	}
	resSuccess(ctx, wrUsers)
}
// @Tags wrUsers表操作
// @Summary 【update】根据id更新数据
// @Description 根据id更新数据
// @Accept x-www-form-urlencoded
// @Produce json
// @Param   id	body	string 	true	"主键更新依据此id"
// @Success 200 {object} SgrResp
// @Router /wrUsers/update-by-id [put]
func (c * WrUsersController) UpdateById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	wrUsers := new(et.WrUsers)
	getPostStructData(ctx, wrUsers)
	has, err := c.serv.UpdateById(id, wrUsers)
	if err != nil {
		resError(ctx, et.EntityFailure, err.Error())
	}else{
		if has == 0 {
			resError(ctx, et.EntityFailure, "影响行数0")
		}else{
			resSuccess(ctx, gin.H{
				"update_count":has,
			})
		}
	}
}