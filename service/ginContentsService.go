package service

import (
	et "gpi/entities"
	"gpi/models"
)

type GinContentsService struct {
}

/**
 * 根据多条件查询数据
 */
func (c *GinContentsService) Find(conditions *et.GinContents, pagination *et.Pagination) (*et.GinContentsPageDao, error) {
	ginContentsModel := models.GinContentsModel{}
	ginContentsPage, err := ginContentsModel.Find(conditions, pagination)
	if err != nil {
		return nil, err
	}
	return ginContentsPage, nil
}

func (c *GinContentsService) FindById(id int) (*et.GinContents, error) {
	ginContentsModel := models.GinContentsModel{}
	return ginContentsModel.GetById(id)
}

func (c *GinContentsService) Insert(ginContents *et.GinContents) (err error) {
	ginContentsModel := models.GinContentsModel{}
	err = ginContentsModel.Insert(ginContents)
	if err != nil {
		return err
	}
	return nil
}

func (c *GinContentsService) UpdateById(id int, ginContents *et.GinContents) (has int64, err error) {
	ginContentsModel := models.GinContentsModel{}
	has, err = ginContentsModel.UpdateById(id, ginContents)
	return
}
