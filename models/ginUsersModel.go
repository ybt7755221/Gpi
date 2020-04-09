package models

import (
	"errors"
	"fmt"
	. "gpi/entities"
	DB "gpi/libraries/database"
	"strings"
)

type GinUsersModel struct {
}

//查找多条数据
func (u *GinUsersModel) Find(conditions *GinUsers, pagination *Pagination) (*GinUsersPageDao, error) {
	dbConn := DB.GetDB(Gin)
	dbConn.ShowSQL(true)
	defer dbConn.Close()
	//获取分页信息
	pageinfo := getPagingParams(pagination)
	limit := pageinfo["limit"].(int)
	offset := pageinfo["offset"].(int)
	dbC := dbConn.Limit(limit, offset)
	defer dbC.Close()
	ginUsersPage := new(GinUsersPageDao)
	fmt.Println(pagination)
	ginUsersPage.PageNum = pagination.PageNum
	ginUsersPage.PageSize = pagination.PageSize
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
	err := dbC.Find(&ginUsersPage.List, conditions)
	total, err := dbC.Count(conditions)
	if err == nil {
		ginUsersPage.Total = total
	}
	return ginUsersPage, err
}

//根据id查找单条数据
func (u *GinUsersModel) GetById(id int) (*GinUsers, error) {
	fmt.Println(id)
	ginUsers := &GinUsers{Id: id}
	dbConn := DB.GetDB(Gin)
	_, err := dbConn.Get(ginUsers)
	defer dbConn.Close()
	return ginUsers, err
}

//插入
func (u *GinUsersModel) Insert(ginUsers *GinUsers) (err error) {
	fmt.Println(ginUsers)
	dbConn := DB.GetDB(Gin)
	affected, err := dbConn.Insert(ginUsers)
	if err != nil {
		return err
	}
	defer dbConn.Close()
	if affected < 1 {
		err = errors.New("插入影响行数: 0")
		return err
	}
	return err
}

//根据id更新
func (u *GinUsersModel) UpdateById(id int, ginUsers *GinUsers) (affected int64, err error) {
	dbConn := DB.GetDB(Gin)
	affected, err = dbConn.Id(id).Update(ginUsers)
	defer dbConn.Close()
	return
}
