package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"gpi/libriries/config"
	"strconv"
	"sync"
)

var Engine *xorm.Engine
var once sync.Once

//自动加载mysql连接
func init() {
	connect()
}

//连接数据库--单例模式
func connect() {
	once.Do(func() {
		var err error
		conf := config.GetSectionMapString("database")
		if err != nil {
			fmt.Println("connect db Error: ", err.Error())
		}
		addrStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
			conf["user"],
			conf["passwd"],
			conf["host"],
			conf["port"],
			conf["name"],
			conf["charset"],
		)
		Engine, err = xorm.NewEngine(conf["driver"], addrStr)
		openMaxInt, _ := strconv.Atoi(conf["openMax"])
		idleMaxInt, _ := strconv.Atoi(conf["idleMax"])
		Engine.SetMaxOpenConns(openMaxInt)
		Engine.SetMaxIdleConns(idleMaxInt)
		if err != nil {
			fmt.Println("Connect DB Error :", err.Error())
		} else {
			fmt.Println("Connect DB Success!")
		}
	})
}
