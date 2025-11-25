package repositories

import (
	"github.com/iamahless/product-filter-api/internal/dtos/filters"
	"github.com/iamahless/product-filter-api/internal/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindByFilters(filters filters.ProductFilter) ([]models.Product, error)
}

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{DB: db}
}

func (repo *productRepository) FindByFilters(filters filters.ProductFilter) ([]models.Product, error) {
	var products []models.Product

	query := repo.DB.Model(&models.Product{})

	if filters.Category != "" {
		query = query.Where("category = ?", filters.Category)
	}

	if filters.PriceLessThan > 0 {
		query = query.Where("price <= ?", filters.PriceLessThan)
	}

	if filters.Limit > 0 {
		page := max(filters.Page, 1)

		offset := (page - 1) * filters.Limit
		query = query.Limit(filters.Limit).Offset(offset)
	}

	err := query.Find(&products).Error

	return products, err
}
