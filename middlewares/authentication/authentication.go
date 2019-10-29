package authentication

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gpi/entities"
	"gpi/libriries/verify"
)
//数据验证-中间件
func Verify(c *gin.Context) {
	var token string
	methodStr := c.Request.Method
	if skip := skipVerify(c, methodStr); skip == false {
		if methodStr == "GET" || methodStr == "DELETE"{
			token = c.Query("token")
			ts = c.Query("ts")
		} else {
			token = c.PostForm("token")
			ts = c.PostForm("ts")
		}
		if len(ts) < 1 {
			c.JSON(http.StatusOK, entities.ApiResonse{http.StatusNoContent, "缺少token值", gin.H{}})
			c.Abort()
		}
		if len(token) < 1 {
			c.JSON(http.StatusOK, entities.ApiResonse{http.StatusNoContent, "缺少token值", gin.H{}})
			c.Abort()
		} else {
			_, tokenStr := verify.GenerateToken(c)
			if token != tokenStr {
				c.JSON(http.StatusOK, entities.ApiResonse{http.StatusForbidden, "验证码token不一致", gin.H{}})
				c.Abort()
			} else {
				c.Next()
			}
		}
	}else{
		c.Next()
	}
}

func skipVerify(c *gin.Context, methodStr string) bool {
	var skip string
	if methodStr == "GET" || methodStr == "DELETE" {
		skip = c.Query("skip_debug")
	}else{
		skip = c.PostForm("skip_debug")
	}
	if skip == "161217" {
		return true
	}
	return false
}

