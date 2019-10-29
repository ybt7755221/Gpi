package timeout

import (
	"context"
	"github.com/gin-gonic/gin"
	"gpi/libriries/config"
	"net/http"
	"time"
)

func Done() gin.HandlerFunc {
	return func(c *gin.Context) {
		conf := config.Config{}
		conf.LoadYamlConfig("sys")
		duration, _ := conf.GetTimeDuration("duration")
		ctx, cancel := context.WithTimeout(c.Request.Context(), duration*time.Second)
		defer func() {
			if ctx.Err() == context.DeadlineExceeded {
				c.JSON(http.StatusGatewayTimeout, "本次请求超时")
				c.Abort()
			}
			cancel()
		}()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func TimeHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		conf := config.Config{}
		conf.LoadYamlConfig("sys")
		duration, _ := conf.GetTimeDuration("duration")
		ctx := c.Request.Context()

		type responseData struct {
			status int
			body   map[string]interface{}
		}

		doneChan := make(chan responseData)
		go func() {
			time.Sleep(duration*time.Second)
			doneChan <- responseData{
				status: 200,
				body:   gin.H{"msg": "请求超时"},
			}
		}()

		select {
			case <-ctx.Done():
				return
			case res := <-doneChan:
				c.JSON(res.status, res.body)
		}
	}
}
