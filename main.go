package main

import (
	"fmt"
	"gpi/libriries/config"
	"gpi/router"
	"log"
	"net/http"
	"time"
)

func main() {
	routers := router.InitRouter()
	sysconf := config.Config{}
	sysconf.LoadYamlConfig("sys")
	duration, _ := sysconf.GetInt64("duration")
	serv := &http.Server{
		Addr: sysconf.GetString("port"),
		Handler:routers,
		ReadTimeout: time.Duration(duration)*time.Second,
		WriteTimeout: 2*time.Duration(duration)*time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := serv.ListenAndServe(); err != nil {
		log.Println(err.Error())
	}else {
		fmt.Println("The Server Listen Port is 9090")
	}
}