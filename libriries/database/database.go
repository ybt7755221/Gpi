package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"gpi/libriries/config"
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
		conf := config.Config{}
		err = conf.LoadYamlConfig("database")
		if err != nil {
			fmt.Println("connect db Error: ", err.Error())
		}
		addrStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
			conf.GetString("user"),
			conf.GetString("passwd"),
			conf.GetString("host"),
			conf.GetString("port"),
			conf.GetString("name"),
			conf.GetString("charset"),
		)
		Engine, err = xorm.NewEngine(conf.GetString("driver"), addrStr)
		if err != nil {
			fmt.Println("Connect DB Error :", err.Error())
		} else {
			fmt.Println("Connect DB Success!")
		}
	})
}
