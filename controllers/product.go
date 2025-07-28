package controllers

import (
	"github.com/barcode-api/repository"
	"github.com/gin-gonic/gin"
)

type productController struct {
	Repository repository.ProductRepository
}

func NewProductController(repo repository.ProductRepository) *productController {
	return &productController{
		Repository: repo,
	}
}

func (p *productController) Create(c *gin.Context) {

}
