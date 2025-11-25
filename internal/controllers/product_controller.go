package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamahless/product-filter-api/internal/dtos/filters"
	"github.com/iamahless/product-filter-api/internal/services"
)

type ProductController struct {
	Service *services.ProductService
}

func NewProductController(service *services.ProductService) *ProductController {
	return &ProductController{Service: service}
}

func (c *ProductController) ListProducts(context *gin.Context) {
	filters := filters.NewProductFilter()
	context.ShouldBindQuery(&filters)

	productResponse, err := c.Service.ListProducts(filters)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, productResponse)
}

func (c *ProductController) GetProduct(context *gin.Context) {
	context.JSON(200, gin.H{"message": "GetProduct endpoint"})
}
