package redis

import (
	"github.com/go-redis/redis"
	"gpi/libriries/config"
	"sync"
)

var Once sync.Once
var Cli *redis.Client

func init()  {
	connect()
}

func connect() {
	Once.Do(func() {
		conf := config.Config{}
		conf.LoadYamlConfig("redis")
		Cli = redis.NewClient(&redis.Options{
			Addr:     conf.GetString("host") + ":" + conf.GetString("port"),
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	})
}