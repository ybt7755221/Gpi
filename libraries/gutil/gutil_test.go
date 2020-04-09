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

func TestFirstToLower(t *testing.T) {
	s := FirstToLower("Hello World")
	t.Log(s)
}

func TestFirstToUpper(t *testing.T) {
	s := FirstToUpper("helloWorld")
	t.Log(s)
}

func BenchmarkBeanUtil(b *testing.B) {
	usersDao := new(dao.GinUsersDao)
	users := new(entities.GinUsers)
	users.Id = 100
	users.Mobile = "11111111111"
	users.Fullname = "fullame"
	users.Email = "asdfoas@dsfs.com"
	users.Username = "sadfas"
	users.Password = "sdfasdfds"
	users.CreateTime = time.Now()
	users.UpdateTime = time.Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BeanUtil(usersDao, users)
	}
}

func BenchmarkTwoJson(b *testing.B) {
	usersDao := new(dao.GinUsersDao)
	users := new(entities.GinUsers)
	users.Id = 100
	users.Mobile = "11111111111"
	users.Fullname = "fullame"
	users.Email = "asdfoas@dsfs.com"
	users.Username = "sadfas"
	users.Password = "sdfasdfds"
	users.CreateTime = time.Now()
	users.UpdateTime = time.Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoJson(usersDao, users)
	}
}
