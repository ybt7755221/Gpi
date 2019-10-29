package verify

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/url"
	"gpi/libriries/config"
)

func GenerateToken(c *gin.Context) (string, string) {
	var rawStr string
	conf := config.Config{}
	err := conf.LoadYamlConfig("auth")
	if err != nil {
		fmt.Println("GenerateToken Err: ", err.Error())
	}
	methodStr := c.Request.Method
	if methodStr == "GET" {
		rawStr = getParams(c.Request.URL.Query())
	}else{
		rawStr = getParams(c.Request.PostForm)
	}
	rawStr = fmt.Sprintf("%s_%s_%s", conf.GetString("app_id"), rawStr, conf.GetString("secret"))
	fmt.Println("rawStr -- ", rawStr)
	return rawStr, GenerateMD5(rawStr, 32)
}

//处理参数
func getParams(data url.Values) string {
	delete(data, "app_id")
	delete(data, "token")
	if(len(data) > 0) {
		var paramsStr string
		for key, val := range data {
			if key != "token" && key != "app_id" {
				fmt.Println("key - val :", key, val)
				paramsStr += key + "-" + val[0] + "-"
			}
		}
		return paramsStr[0 : len(paramsStr)-1]
	}else{
		return ""
	}
}
//获取md5
func GenerateMD5(raw string, size int) string {
	md5H := md5.New()
	md5H.Write([]byte(raw))
	token := hex.EncodeToString(md5H.Sum(nil))
	if size == 16 {
		return token[8:16]
	}
	return token
}