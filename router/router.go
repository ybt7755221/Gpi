package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	. "gpi/controllers"
	"gpi/middlewares/authentication"
	"gpi/middlewares/exception"
	_ "gpi/docs"
)

func InitRouter() *gin.Engine {
	cts := ContentsController{}
	users := UsersController{}
	demo := DemoController{}
	router := gin.Default()
	router.Use(exception.Recover())
	idx := router.Group("/")
	{
		idx.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		idx.GET("/", func(c *gin.Context) {
			c.String(200, "Gpi系统首页")
		})
		idx.Any("/createToken", demo.Token)
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
