package service

import (
	et "gpi/entities"
	"gpi/models"
)

type GinUsersService struct {
}
/**
 * 根据多条件查询数据
 */
func (c *GinUsersService) Find(params map[string]interface{}) ([]et.GinUsers, error) {
	ginUsersModel := models.GinUsersModel{}
	ginUsersList, err := ginUsersModel.Find(params)
	if err != nil {
		return nil, err
	}
	return ginUsersList, nil
}

func (c *GinUsersService) FindById(id int) (*et.GinUsers, error) {
	ginUsersModel := models.GinUsersModel{}
	return ginUsersModel.GetById(id)
}

func (c *GinUsersService) Insert(ginUsers *et.GinUsers) (err error) {
	ginUsersModel := models.GinUsersModel{}
	err = ginUsersModel.Insert(ginUsers)
	if err != nil {
		return err
	}
	return nil
}

func (c *GinUsersService) UpdateById(id int, ginUsers *et.GinUsers) (has int64, err error) {
	ginUsersModel := models.GinUsersModel{}
	has, err = ginUsersModel.UpdateById(id, ginUsers)
	return
}