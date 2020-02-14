package main

import (
	"fmt"
	"gpi/config"
	"gpi/libraries/elog"
	"gpi/router"
	"net/http"
	"time"
)

func main() {
	routers := router.InitRouter()
	duration := config.Duration
	port := config.HttpPort
	serv := &http.Server{
		Addr:           ":" + port,
		Handler:        routers,
		ReadTimeout:    time.Duration(duration) * time.Second,
		WriteTimeout:   2 * time.Duration(duration) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := serv.ListenAndServe(); err != nil {
		elog.ErrMail(err.Error(), elog.GetFileInfo())
	} else {
		fmt.Println("The Server Listen Port is " + port)
	}
}
