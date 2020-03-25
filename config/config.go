package config

import (
	"gpi/libraries/apolloCli"
	"strconv"
)

const (
	AppName 	  = "Gpi"
	Secret 		  = "Dl*sCKW7C{SfYiPtYX*O5/71vG9&sm?2U"
	HttpPort 	  = "9090"
	Duration 	  = "TIMEOUT"
	LogPath  	  = "/data/logs"
	WechatUrl	  = "WECHAT"
	WechattoUser  = "WECHAT_TOUSER"
	WechatSecret  = "WECHAT_SECRET"
	WechatAppkey  = "WECHAT_APPKEY"
	KafkaUrl	  = "KFKURL"
	KafKaProt	  = "KFKPORT"
)

type JaegerConf struct {
	Host string
	Port string
	Type string
}

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