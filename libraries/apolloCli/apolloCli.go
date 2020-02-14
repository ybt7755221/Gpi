package apolloCli

import (
	"fmt"
	"github.com/shima-park/agollo"
)

const (
	AppId = "Gpi"
	Ip = "127.0.0.1"
	NameSpacename = "application"
	BackUpFile = "/Users/Burt/Work/logs/application.agollo"
)

func OptionInit() map[string]interface{} {
	apoCli, err := agollo.New(Ip, AppId,
		agollo.BackupFile(BackUpFile),
		agollo.FailTolerantOnBackupExists(),
		agollo.AutoFetchOnCacheMiss(),
	)
	if err != nil {
		panic(err.Error())
	}else{
		fmt.Println("apollo start is successd")
	}
	confMap := apoCli.GetNameSpace(NameSpacename)
	apoCli.Start()  // Start后会启动goroutine监听变化，
	go func() {
		watchCh := apoCli.Watch()
		for{
			select{
			case resp := <-watchCh:
				txtFile := "【Apollo Update】Apollo has modified！Namespace is "+resp.Namespace
				fmt.Println(txtFile)
			}
		}
	}()
	return confMap
}