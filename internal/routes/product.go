package routes

import (
	"github.com/barcode-api/controllers"
	"github.com/barcode-api/internal/database"
	"github.com/barcode-api/middleware/auth"
	"github.com/barcode-api/repository"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine, jwtSecret string) {

	db := database.GetDb()
	productRepository := repository.NewProductRepository(db)
	productController := controllers.NewProductController(productRepository)

	group := router.Group("/product")
	group.Use(auth.Authentication(jwtSecret))
	group.POST("", productController.Create)
}
