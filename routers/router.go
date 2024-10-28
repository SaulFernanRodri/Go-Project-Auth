package routers

import (
	"auth/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(authController *controllers.AuthController) *gin.Engine {

	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}

	return router
}
