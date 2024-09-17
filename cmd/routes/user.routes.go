package routes

import (
	"websocket/cmd/handlers"
	"websocket/internal/service"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.RouterGroup, userSrv service.UserService) {
	handler := handlers.NewUserHandler(userSrv)
	user := router.Group("/users")
	{
		user.GET("/:userId", handler.Get)
		user.GET("", handler.GetAll)
	}

	auth := router.Group("/auth")
	{
		auth.POST("/register", handler.Create)
		auth.POST("/login", handler.Login)
	}
}
