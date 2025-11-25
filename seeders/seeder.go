package seeders

import (
	"log"

	"github.com/iamahless/product-filter-api/internal/models"
	"gorm.io/gorm"
)

func Load(db *gorm.DB) {
	products := []models.Product{
		{Name: "BV Lean leather ankle boots", Sku: "000001", Category: "boots", Price: 89000},
		{Name: "BV Lean leather ankle boots", Sku: "000002", Category: "boots", Price: 99000},
		{Name: "Ashlington leather ankle boots", Sku: "000003", Category: "boots", Price: 71000},
		{Name: "Naima embellished suede sandals", Sku: "000004", Category: "sandals", Price: 79500},
		{Name: "Nathane leather sneakers", Sku: "000005", Category: "sneaker", Price: 59000},
	}

	for _, product := range products {
		if err := db.FirstOrCreate(&product, models.Product{Sku: product.Sku}).Error; err != nil {
			log.Printf("Could not seed user %s: %v", product.Sku, err)
		}
	}
	log.Println("Database seeded successfully!")
}
