package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strconv"
	"time"
)

type Config struct {
	v *viper.Viper
	ConfMap map[string]string
}

func (c *Config) LoadYamlConfig(section string) error {
	//判断测试还是正式
	var fileName string
	ENV := os.Getenv("GPIACTIVE")
	if (ENV == "pro") {
		fileName = "config"
	}else{
		fileName = "config_fat"
	}
	c.v = viper.New()
	c.v.SetConfigName(fileName)
	c.v.AddConfigPath("config/")
	c.v.SetConfigType("yaml")
	if err := c.v.ReadInConfig(); err != nil {
		fmt.Printf("err:%s\n",err)
		return err
	}
	c.ConfMap = c.v.GetStringMapString(section)
	fmt.Println("config: " + fileName)
	return nil
}

func (c *Config) GetString(field string) string {
	return c.ConfMap[field]
}

func (c *Config) GetInt(field string) (int, error) {
	res := c.GetString(field)
	fieldInt, err := strconv.Atoi(res)
	return fieldInt, err
}

func (c *Config) GetInt64(field string) (int64, error) {
	res := c.GetString(field)
	fieldInt64, err := strconv.ParseInt(res, 10, 64)
	return fieldInt64, err
}

func (c *Config) GetTimeDuration(field string) (time.Duration, error) {
	fieldInt64, err := c.GetInt64(field)
	duration := time.Duration(fieldInt64)
	return duration, err
}

