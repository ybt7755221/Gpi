package main

import (
	conf "gpi/config"
	"gpi/libraries/elog"
	"gpi/router"
	"fmt"
	"github.com/DeanThompson/ginpprof"
	"net/http"
	_ "net/http/pprof"
	"strings"
	"time"
)
// @title gpi文档平台
// @version 1.0
// @description gpi自动文档
// @host localhost
// @BasePath /
func main() {
	//初始化router
	routers := router.InitRouter()
	//正式环境关闭prof
	pprofStauts := conf.GetApolloString("PPROF_STATUS", "stop")
	if strings.ToLower(pprofStauts) == "start" {
		ginpprof.Wrap(routers)
	}
	duration := conf.GetApolloInt(conf.Duration, 30)
	//初始化服务
	serv := &http.Server{
		Addr: ":" + conf.HttpPort,
		Handler:routers,
		ReadTimeout: time.Duration(duration)*time.Second,
		WriteTimeout: 2*time.Duration(duration)*time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := serv.ListenAndServe(); err != nil {
		elog.New(err.Error(), elog.GetFileInfo(1))
	}else {
		fmt.Println("The Server Listen Port is 8080")
	}
}