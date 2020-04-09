package models

import (
	"errors"
	"fmt"
	. "gpi/entities"
	DB "gpi/libraries/database"
	"strings"
)

type GinContentsModel struct {
}

//查找多条数据
func (u *GinContentsModel) Find(conditions *GinContents, pagination *Pagination) (*GinContentsPageDao, error) {
	dbConn := DB.GetDB(Gin)
	defer dbConn.Close()
	//获取分页信息
	pageinfo := getPagingParams(pagination)
	limit := pageinfo["limit"].(int)
	offset := pageinfo["offset"].(int)
	dbC := dbConn.Limit(limit, offset)
	defer dbC.Close()
	ginContentsPage := new(GinContentsPageDao)
	ginContentsPage.PageNum = pagination.PageNum
	ginContentsPage.PageSize = pagination.PageSize
	//排序
	sort := pageinfo["sort"].(map[string]string)
	if len(sort) > 0 {
		for key, val := range sort {
			if strings.ToLower(val) == "asc" {
				dbC = dbC.Asc(key)
			} else {
				dbC = dbC.Desc(key)
			}
		}
	}
	//执行查找
	err := dbC.Find(&ginContentsPage.List, conditions)
	total, err := dbC.Count(conditions)
	if err == nil {
		ginContentsPage.Total = total
	}
	return ginContentsPage, err
}

//根据id查找单条数据
func (u *GinContentsModel) GetById(id int) (*GinContents, error) {
	fmt.Println(id)
	ginContents := &GinContents{Id: id}
	dbConn := DB.GetDB(Gin)
	_, err := dbConn.Get(ginContents)
	defer dbConn.Close()
	return ginContents, err
}

//插入
func (u *GinContentsModel) Insert(ginContents *GinContents) (err error) {
	dbConn := DB.GetDB(Gin)
	affected, err := dbConn.Insert(ginContents)
	defer dbConn.Close()
	if err != nil {
		return err
	}
	if affected < 1 {
		err = errors.New("插入影响行数: 0")
		return err
	}
	return err
}

//根据id更新
func (u *GinContentsModel) UpdateById(id int, ginContents *GinContents) (affected int64, err error) {
	dbConn := DB.GetDB(Gin)
	affected, err = dbConn.Id(id).Update(ginContents)
	defer dbConn.Close()
	return
}
