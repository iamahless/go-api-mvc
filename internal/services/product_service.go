package services

import (
	"github.com/iamahless/product-filter-api/internal/dtos"
	"github.com/iamahless/product-filter-api/internal/dtos/filters"
	"github.com/iamahless/product-filter-api/internal/models"
	"github.com/iamahless/product-filter-api/internal/repositories"
)

type ProductService struct {
	Repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) *ProductService {
	return &ProductService{Repo: repo}
}

func (s *ProductService) ListProducts(filters filters.ProductFilter) ([]dtos.ProductResponse, error) {
	products := []models.Product{}
	productResponses := []dtos.ProductResponse{}

	products, err := s.Repo.FindByFilters(filters)

	if err != nil {
		return nil, err
	}

	for _, product := range products {
		priceResponse := dtos.NewProductPriceResponse(product.Price, product.Price, "10%")
		productResponse := dtos.ProductResponse{
			Sku:      product.Sku,
			Name:     product.Name,
			Category: product.Category,
			Price:    priceResponse,
		}
		productResponses = append(productResponses, productResponse)
	}

	return productResponses, nil
}
