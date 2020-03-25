package router

import (
	"gpi/config"
	. "gpi/controllers"
	_ "gpi/docs"
	"gpi/middlewares/exception"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	"time"
)

func InitRouter() *gin.Engine {
	demo := DemoController{}
	router := gin.Default()
	router.Use(exception.Recover())
	//访问日志吸入文件样式
	fileDir := config.LogPath+string(os.PathSeparator)+config.AppName+"_gin_access_"+time.Now().Format("2006-01-02")+".log"
	file, _ := os.Create(fileDir)
	c := gin.LoggerConfig{
		Output:file,
		SkipPaths:[]string{"/test"},
		Formatter: func(params gin.LogFormatterParams) string {
			return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
				params.ClientIP,
				params.TimeStamp.Format(time.RFC1123),
				params.Method,
				params.Path,
				params.Request.Proto,
				params.StatusCode,
				params.Latency,
				params.Request.UserAgent(),
				params.ErrorMessage,
			)
		},
	}
	router.Use(gin.LoggerWithConfig(c))
	//各种工具以及需要验签的部分
	idx :=router.Group("/")
	{
		ENVIR := os.Getenv("ACTIVE")
		if ENVIR != "pro" || ENVIR != "uat" {
			//swagger-doc路由
			idx.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
		idx.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"title" : "GPI首页 ",
				"env"  : os.Getenv("ACTIVE"),
				"apollo" : config.GetApolloString("ENVIRONMENT", ""),
			})
		})
		//docker健康检测用，切勿删除
		idx.GET("/health", func(c *gin.Context) {
			c.String(200, "系统-健康检查！要保持畅通！")
		})
		//各种demo，后期可删
		idx.GET("/redis", demo.Redis)
		idx.GET("/email", demo.Email)
		idx.GET("/conf", demo.GetConf)
		idx.GET("/mgo", demo.Mgo)
	}
	//工具库router
	toolRouter(router)
	ginUsersRouter(router)
	ginContentsRouter(router)
	return router
}

func ginUsersRouter(router *gin.Engine) {
	ginUsers := GinUsersController{}
	ginUsersR := router.Group("users")
	{
		ginUsersR.GET("/", ginUsers.Find)
		ginUsersR.POST("/", ginUsers.Create)
		ginUsersR.GET("/get-by-id/:id", ginUsers.FindById)
		ginUsersR.POST("/update-by-id", ginUsers.UpdateById)
	}
}

func ginContentsRouter(router *gin.Engine) {
	ginContents := GinContentsController{}
	ginContentsR := router.Group("contents")
	{
		ginContentsR.GET("/", ginContents.Find)
		ginContentsR.POST("/", ginContents.Create)
		ginContentsR.GET("/get-by-id/:id", ginContents.FindById)
		ginContentsR.POST("/update-by-id", ginContents.UpdateById)
	}
}

func toolRouter(router *gin.Engine) {
	toolC := ToolController{}
	//工具库
	toolR := router.Group("tool")
	{
		toolR.Any("/get-token", toolC.Token)
		toolR.POST("/get-jwt-token", toolC.GetJWTToken)
		toolR.POST("/parse-jwt-token", toolC.ParseJWTToken)
	}
	kfkR := router.Group("kafka")
	{
		kfkR.POST("/", toolC.SendKafka)
	}
}
