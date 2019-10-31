package main

import (
	"fmt"
	"gpi/libriries/config"
	"gpi/libriries/elog"
	"gpi/router"
	"net/http"
	"time"
)

func main() {
	routers := router.InitRouter()
	duration := config.Conf.GetInt64("sys.duration")
	port := config.Conf.GetString("sys.port")
	serv := &http.Server{
		Addr: ":"+port,
		Handler:routers,
		ReadTimeout: time.Duration(duration)*time.Second,
		WriteTimeout: 2*time.Duration(duration)*time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := serv.ListenAndServe(); err != nil {
		elog.ErrMail(nil, err.Error())
	}else {
		fmt.Println("The Server Listen Port is " + port)
	}
}