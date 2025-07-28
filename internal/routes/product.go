package routes

import (
	"github.com/barcode-api/controllers"
	"github.com/barcode-api/db"
	"github.com/barcode-api/repository"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {

	db := db.GetDb()
	productRepository := repository.NewProductRepository(db)
	productController := controllers.NewProductController(productRepository)

	router.POST("/product", productController.Create)
}
