package main

import (
	"flag"

	"github.com/iamahless/product-filter-api/config"
	"github.com/iamahless/product-filter-api/internal/controllers"
	"github.com/iamahless/product-filter-api/internal/models"
	"github.com/iamahless/product-filter-api/internal/routes"
	"github.com/iamahless/product-filter-api/internal/services"
	"github.com/iamahless/product-filter-api/seeders"
)

func main() {
	db := config.SetupDB()

	seed := flag.Bool("seed", false, "Run database seeders")
	migrate := flag.Bool("migrate", false, "Run database migrations")
	flag.Parse()

	if *migrate {
		db.AutoMigrate(&models.Product{})
	}

	if *seed {
		seeders.Load(db)
		return
	}

	productService := services.NewProductService(db)
	productController := controllers.NewProductController(productService)

	router := routes.SetupRouter(productController)

	router.Run(":8080")
}
