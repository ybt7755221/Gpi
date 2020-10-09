package router

import (
	. "gpi/controllers"
	"github.com/gin-gonic/gin"
)

func wrUsersRouter(router *gin.Engine) {
	wrUsers := WrUsersController{}
	wrUsersR := router.Group("users")
	{
		wrUsersR.GET("/", wrUsers.Find)
		wrUsersR.GET("/page", wrUsers.FindPaging)
		wrUsersR.POST("/", wrUsers.Create)
		wrUsersR.GET("/find-by-id/:id", wrUsers.FindById)
		wrUsersR.PUT("/update-by-id", wrUsers.UpdateById)
	}
}