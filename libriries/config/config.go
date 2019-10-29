package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"sync"
)

var Conf *viper.Viper
var once sync.Once

func init() {
	err := loadYamlFile()
	if err != nil {
		panic(err)
	}
}

func loadYamlFile() error {
	var err error
	//判断测试还是正式
	once.Do(func() {
		var fileName string
		ENV := os.Getenv("GPIACTIVE")
		if (ENV == "pro") {
			fileName = "config"
		}else{
			fileName = "config_fat"
		}
		Conf = viper.New()
		Conf.SetConfigName(fileName)
		Conf.AddConfigPath("config/")
		Conf.SetConfigType("yaml")
		if err = Conf.ReadInConfig(); err != nil {
			fmt.Printf("err:%s\n",err)
		}
	})
	return err
}

func GetSectionMap(section string) map[string]interface{}{
	return Conf.GetStringMap(section)
}

func GetSectionMapString(section string) map[string]string{
	return Conf.GetStringMapString(section)
}
