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
func (c *GinUsersService) Find(conditions *et.GinUsers, pagination *et.Pagination) (*et.GinUsersPageDao, error) {
	ginUsersModel := models.GinUsersModel{}
	ginUsersPage, err := ginUsersModel.Find(conditions, pagination)
	if err != nil {
		return nil, err
	}
	return ginUsersPage, nil
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
