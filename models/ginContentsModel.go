package models

import (
	. "gpi/entities"
	DB "gpi/libraries/database"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type GinContentsModel struct {
}
//查找多条数据
func (u *GinContentsModel) Find(params map[string]interface{}) ([]GinContents, error) {
	dbConn := DB.GetDB(Gin)
	defer dbConn.Close()
	ginContents := make([]GinContents, 0)
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
	err := dbC.Find(&ginContents)
	return ginContents, err
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
	if affected < 1 {
		err = errors.New("插入影响行数: 0" )
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