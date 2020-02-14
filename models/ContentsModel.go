package models

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	. "gpi/entities"
	DB "gpi/libraries/database"
	"strings"
	"time"
)

type ContentsModel struct {
}

/**
 * 获取用户信息
 * @Param params 请求参数
 * @return {object} []GinContents, {object} error
 */
func (u *ContentsModel) GetContents(params gin.H) ([]GinContents, error) {
	Contents := make([]GinContents, 0)
	dbConn := DB.GetDB(Gin)
	dbC := dbConn.Where("1")
	for key, val := range params["conditions"].(map[string]string) {
		if len(val) > 0 {
			if key == "topic" {
				dbC = dbC.And(key+" LIKE ?", "%"+val+"%")
			} else {
				dbC = dbC.And(key+" = ?", val)
			}
		}
	}
	dbC = dbC.Limit(params["limit"].(int), params["offset"].(int))
	//排序
	sort := params["sort"].(map[string]string)
	if len(sort) > 0 {
		for key, val := range sort{
			if strings.ToLower(val) == "asc" {
				dbC = dbC.Asc(key)
			}else{
				dbC = dbC.Desc(key)
			}
		}
	}
	err := dbC.Find(&Contents)
	return Contents, err
}

/**
 * 根据Id获取用户信息
 * @Param params 请求参数
 * @Param offset 起始
 * @Param limit  长度
 * @return {object} GinContents, {object} error
 */
func (u *ContentsModel) GetById(id int) (*GinContents, error) {
	user := &GinContents{Id: id}
	dbConn := DB.GetDB(Gin)
	_, err := dbConn.Get(user)
	return user, err
}

func (u *ContentsModel) Insert(conn *GinContents) (err error) {
	if conn.Topic == "" {
		err = errors.New("topic不能为空!")
		return err
	}
	conn.OpTime = time.Now()
	dbConn := DB.GetDB(Gin)
	affected, err := dbConn.Insert(conn)
	if err == nil && affected < 1 {
		err = errors.New("插入影响行数: 0")
	}
	return err
}

func (u *ContentsModel) UpdateById(id int, conn *GinContents) (affected int64, err error) {
	conn.OpTime = time.Now()
	dbConn := DB.GetDB(Gin)
	affected, err = dbConn.Id(id).Update(conn)
	return
}

func (u *ContentsModel) Delete(id int) error {
	conn := new(GinContents)
	dbConn := DB.GetDB(Gin)
	affected, err := dbConn.Id(id).Delete(conn)
	if err == nil && affected < 1 {
		err = errors.New("插入影响行数: 0")
	}
	return err
}
