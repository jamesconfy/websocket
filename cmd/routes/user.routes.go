package routes

import (
	"project-name/cmd/handlers"
	"project-name/internal/service"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.RouterGroup, userSrv service.UserService) {
	handler := handlers.NewUserHandler(userSrv)
	user := router.Group("/users")
	{
		user.POST("", handler.Create)
		user.POST("/login", handler.Login)
	}
}
