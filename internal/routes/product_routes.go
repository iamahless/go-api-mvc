package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iamahless/product-filter-api/internal/controllers"
)

func RegisterProductRoutes(router *gin.RouterGroup, pc *controllers.ProductController) {
	products := router.Group("/products")
	{
		products.GET("/", pc.ListProducts)
		products.GET("/:id", pc.GetProduct)
	}
}
