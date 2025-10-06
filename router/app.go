package router

import (
	"flashchat/service"

	"flashchat/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/index", service.GetIndex)
	r.GET("/users/getUsers", service.GetUsers)
	r.POST("/users/createUser", service.CreateUser)
	r.POST("/users/updateUser", service.UpdateUser)
	r.DELETE("/users/deleteUser", service.DeleteUser)

	// message
	r.GET("/chat/sendMsg", service.SendMsg)
	r.GET("/chat/sendUserMsg", service.SendUserMsg)
	return r
}
