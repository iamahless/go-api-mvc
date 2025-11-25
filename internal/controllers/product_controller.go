package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/iamahless/product-filter-api/internal/services"
)

type ProductController struct {
	Service *services.ProductService
}

func NewProductController(service *services.ProductService) *ProductController {
	return &ProductController{Service: service}
}

func (c *ProductController) ListProducts(context *gin.Context) {
	context.JSON(200, gin.H{"message": "ListProducts endpoint"})
}

func (c *ProductController) GetProduct(context *gin.Context) {
	context.JSON(200, gin.H{"message": "GetProduct endpoint"})
}
