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
		conf := config.GetSectionMapString("redis")
		Cli = redis.NewClient(&redis.Options{
			Addr:     conf["host"] + ":" + conf["port"],
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	})
}