package config

import (
	"gpi/libraries/apolloCli"
	"strconv"
)

const (
	AppName 	= "Gpi"
	Secret 		= "Dl*sCKW7C{SfYiPtYX*O5/71vG9&sm?2U"
	HttpPort 	= "9090"
	Duration 	= 2
	LogPath  	= "/Users/Burt/Work/logs"
	PanicPath 	= "/Users/Burt/Work/logs"
	EmailHost 	= "smtp.163.com"
	EmailPort 	= "25"
	EmailUser 	= "burt_yu@163.com"
	EmailPasswd	= "XXXXXXXX"
	EmailTo		= "ybt7755221@sohu.com"
	EmailErrTopic = "【Gpi系统】报错"
	JaegerHost = "127.0.0.1"
	JaegerPort = "6831"
	JaegerType = "const"
)

var ApoCli map[string]interface{}

func init() {
	apolloCli.OptionInit()
}

func GetApolloString(key string, defValue string) string {
	apoCli := apolloCli.GetApolloConfig()
	if apoCli[key] == nil {
		return defValue
	}
	return apoCli[key].(string)
}

func GetApolloInt(key string, defValue int) int {
	value := GetApolloString(key, "")
	if len(value) == 0 {
		return defValue
	}
	num, _ := strconv.Atoi(value)
	return num
}