package models

import (
	. "gpi/entities"
	DB "gpi/libraries/database"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type GinUsersModel struct {
}
//查找多条数据
func (u *GinUsersModel) Find(params map[string]interface{}) ([]GinUsers, error) {
	dbConn := DB.GetDB(Gin)
	defer dbConn.Close()
	ginUsers := make([]GinUsers, 0)
	dbC := dbConn.Where("1")
	defer dbC.Close()
	reflect.TypeOf(params["conditions"])
	//where条件
	conditions := params["conditions"].(map[string]string)
	if len(conditions) > 0 {
		for key, val := range params["conditions"].(map[string]string) {
			if len(val) > 0 {
				dbC = dbC.And(key+" = ?", val)
			}
		}
	}
	//limit
	dbC = dbC.Limit(params["limit"].(int), params["offset"].(int))
	if params["sortField"] == "" {
		params["sortField"] = "id"
	}
	//排序
	sort := params["sort"].(map[string]string)
	fmt.Println(len(sort))
	if len(sort) > 0 {
		for key, val := range sort{
			if strings.ToLower(val) == "asc" {
				dbC = dbC.Asc(key)
			}else{
				dbC = dbC.Desc(key)
			}
		}
	}
	err := dbC.Find(&ginUsers)
	return ginUsers, err
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
		err = errors.New("插入影响行数: 0" )
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