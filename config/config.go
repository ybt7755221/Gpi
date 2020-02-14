package config

import (
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
	EmailPasswd	= "Burtyu1989"
	EmailTo		= "ybt7755221@sohu.com"
	EmailErrTopic = "【Gpi系统】报错"
	JaegerHost = "127.0.0.1"
	JaegerPort = "6831"
	JaegerType = "const"
)

var ApoCli map[string]interface{}

func init() {
	ApoCli = map[string]interface{}{}
	//ApoCli = apolloCli.OptionInit() //如果使用apollo用这个替换上面的
}

func GetApolloString(key string, defValue string) string {
	if ApoCli[key] == nil {
		return defValue
	}
	return ApoCli[key].(string)
}

func GetApolloInt(key string, defValue int) int {
	value := GetApolloString(key, "")
	if len(value) == 0 {
		return defValue
	}
	num, _ := strconv.Atoi(value)
	return num
}