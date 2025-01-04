package route

import (
	"go_auth/controller"
	"go_auth/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(api *gin.RouterGroup) {
	api.POST("/login", middleware.LoginHandler)
	api.GET("/user", controller.GetAllUserHandler)
	api.POST("/user", controller.CreateUserHandler)

	api.Use(middleware.AuthMiddleware())
	{
		api.POST("/logout", middleware.LogoutHandler)
		api.GET("/user/:id", controller.GetUserHandler) // Tambahkan parameter ID
		api.PUT("/user", controller.UpdateUserHandler)
		api.DELETE("/user", controller.DeleteUserHandler)
	}
}