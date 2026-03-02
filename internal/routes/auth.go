package routes

import (
	"github.com/barcode-api/controllers"
	"github.com/barcode-api/internal/database"
	"github.com/barcode-api/repository"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, jwtSecret string) {
	db := database.GetDb()
	userRepository := repository.NewUserRepository(db)
	userController := controllers.NewUserController(&userRepository, jwtSecret)

	group := router.Group("/auth")

	{
		group.POST("/login", userController.Login)
	}
}
