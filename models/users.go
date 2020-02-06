package models

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	. "gpi/entities"
	DB "gpi/libriries/database"
	"gpi/libriries/verify"
	"time"
)

type Users struct {
	Common
}

/**
 * 获取用户信息
 * @Param params 请求参数
 * @return {object} []GinUsers, {object} error
 */
func (u *Users) GetUser(params gin.H) ([]GinUsers, error) {
	dbConn := DB.GetDB()
	defer dbConn.Close()
	users := make([]GinUsers, 0)
	dbC := dbConn.Where("1")
	for key, val := range params["conditions"].(map[string]string) {
		if len(val) > 0 {
			if key == "username" {
				dbC = dbC.And(key+" LIKE ?", "%"+val+"%")
			} else {
				dbC = dbC.And(key+" = ?", val)
			}
		}
	}
	dbC = dbC.Limit(params["limit"].(int), params["offset"].(int))
	if params["sortField"] == "" {
		params["sortField"] = "id"
	}
	if params["sort"].(int) == 1 {
		dbC = dbC.Asc(params["sortField"].(string))
	} else {
		dbC = dbC.Desc(params["sortField"].(string))
	}
	err := dbC.Find(&users)
	return users, err
}

/**
 * 根据Id获取用户信息
 * @Param params 请求参数
 * @Param offset 起始
 * @Param limit  长度
 * @return {object} GinUsers, {object} error
 */
func (u *Users) GetById(id int) (*GinUsers, error) {
	user := &GinUsers{Id: id}
	dbConn := DB.GetDB()
	_, err := dbConn.Get(user)
	defer dbConn.Close()
	return user, err
}

func (u *Users) Insert(user *GinUsers) (err error) {
	if user.Username == "" || user.Password == "" || user.Mobile == "" {
		err = errors.New("username, passowrd, mobile不能为空!")
		return err
	}
	if user.Fullname == "" {
		user.Fullname = user.Username
	}
	if user.Email == "" {
		user.Email = user.Username + "@gpi.com"
	}
	user.Password = verify.GenerateMD5(user.Password, 32)
	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()
	dbConn := DB.GetDB()
	affected, err := dbConn.Insert(user)
	defer dbConn.Close()
	if affected < 1 {
		err = errors.New("插入影响行数: 0")
		return err
	}
	return err
}

func (u *Users) Update(id int, user *GinUsers) (affected int64, err error) {
	if user.Password != "" {
		user.Password = verify.GenerateMD5(user.Password, 32)
	}
	user.UpdateTime = time.Now()
	dbConn := DB.GetDB()
	affected, err = dbConn.Id(id).Update(user)
	defer dbConn.Close()
	return
}
