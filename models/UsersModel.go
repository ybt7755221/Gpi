package models

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	. "gpi/entities"
	DB "gpi/libraries/database"
	"gpi/libraries/verify"
	"strings"
	"time"
)

type UsersModel struct {
}

func (u *UsersModel) GetUser(params gin.H) ([]GinUsers, error) {
	dbConn := DB.GetDB(Gin)
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
	err := dbC.Find(&users)
	return users, err
}

func (u *UsersModel) GetById(id int) (*GinUsers, error) {
	user := &GinUsers{Id: id}
	dbConn := DB.GetDB(Gin)
	_, err := dbConn.Get(user)
	defer dbConn.Close()
	return user, err
}

func (u *UsersModel) Insert(user *GinUsers) (err error) {
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
	dbConn := DB.GetDB(Gin)
	affected, err := dbConn.Insert(user)
	defer dbConn.Close()
	if affected < 1 {
		err = errors.New("插入影响行数: 0")
		return err
	}
	return err
}

func (u *UsersModel) UpdateById(id int, user *GinUsers) (affected int64, err error) {
	if user.Password != "" {
		user.Password = verify.GenerateMD5(user.Password, 32)
	}
	user.UpdateTime = time.Now()
	dbConn := DB.GetDB(Gin)
	affected, err = dbConn.Id(id).Update(user)
	defer dbConn.Close()
	return
}
