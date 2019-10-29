package router

import (
	"github.com/gin-gonic/gin"
	. "gpi/controllers"
	"gpi/middlewares/authentication"
	"gpi/middlewares/exception"
)

func InitRouter() *gin.Engine {
	cts := Contents{}
	users := Users{}
	demo := Demo{}
	router := gin.Default()
	router.Use(exception.Recover())
	idx :=router.Group("/")
	{
		idx.GET("/", func(c *gin.Context) {
			c.String(200, "Gpi系统首页")
		})
		idx.Any("/createToken", users.Token)
		idx.GET("/redis", demo.Redis)
		idx.GET("/email", demo.Email)
		idx.GET("/conf", demo.GetConf)
	}
	ctR := router.Group("/content", authentication.Verify)
	{
		ctR.GET("/", cts.Get)
		ctR.POST("/", cts.Create)
		ctR.GET("/:id", cts.GetId)
		ctR.PUT("/:id", cts.Update)
		ctR.DELETE("/:id", cts.Delete)
	}
	usR := router.Group("/user", authentication.Verify)
	{
		usR.GET("/", users.Get)
		usR.POST("/", users.Create)
		usR.GET("/:id", users.GetId)
		usR.PUT("/:id", users.Update)
	}
	return router
}