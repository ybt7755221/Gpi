package service

import (
	"github.com/gin-gonic/gin"
	"gpi/entities"
	"gpi/models"
)

type ContentsService struct {}

func (c *ContentsService) Find(params gin.H) ([]entities.GinContents, error) {
	contentsModel := models.ContentsModel{}
	return contentsModel.GetContents(params)
}

func (c *ContentsService) FindById(id int) (*entities.GinContents, error) {
	contentsModel := models.ContentsModel{}
	return contentsModel.GetById(id)
}

func (c *ContentsService) Insert(contents *entities.GinContents) error {
	contentsModel := models.ContentsModel{}
	return contentsModel.Insert(contents)
}

func (c *ContentsService) UpdateById(id int, contents *entities.GinContents) (has int64, err error) {
	contentsModel := models.ContentsModel{}
	return contentsModel.UpdateById(id, contents)
}

func (c *ContentsService) DeleteById(id int) (err error) {
	contentsModel := models.ContentsModel{}
	return contentsModel.Delete(id)
}

