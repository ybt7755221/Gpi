package gutil

import (
	"gpi/dao"
	"gpi/entities"
	"testing"
	"time"
)

func TestBeanUtil(t *testing.T) {
	usersDao := new(dao.GinUsersDao)
	users := new(entities.GinUsers)
	users.Id = 100
	users.Mobile = "11111111111"
	users.Fullname = "fullame"
	users.CreateTime = time.Now()
	users.UpdateTime = time.Now()
	BeanUtil(usersDao, users)
	t.Logf("usersDao : %v", usersDao)
}
