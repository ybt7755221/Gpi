package service
import (
	et "gpi/entities"
	"gpi/models"
	"github.com/gin-gonic/gin"
)

type UsersService struct {}

func (u *UsersService) Find(params gin.H) ([]et.GinUsers, error) {
	userModel := models.UsersModel{}
	return userModel.GetUser(params)
}

func (u *UsersService) FindById(id int) (*et.GinUsers, error) {
	userModel := models.UsersModel{}
	return userModel.GetById(id)
}

func (u *UsersService) Insert(users *et.GinUsers) error {
	userModel := models.UsersModel{}
	return userModel.Insert(users)
}

func (u *UsersService) UpdateById(id int, users *et.GinUsers) (has int64, err error) {
	userModel := models.UsersModel{}
	return userModel.UpdateById(id, users)
}