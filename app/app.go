package app

import (
	"auth/config"
	"auth/controllers"
	"auth/internal/database"
	"auth/repositories"
	"auth/routers"
	"auth/services"

	"github.com/gin-gonic/gin"
)

func InitializeApp() *gin.Engine {

	config.LoadConfig()
	db := database.InitDB()
	database.Migrate(db)

	authRepo := repositories.NewAuthRepo(db)
	authService := services.NewAuthService(authRepo)
	authController := controllers.NewAuthController(authService)

	router := routers.SetupRouter(authController)

	return router
}
