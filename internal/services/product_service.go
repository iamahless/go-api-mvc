package services

import (
	"fmt"
	"math"

	"github.com/iamahless/product-filter-api/internal/dtos"
	"github.com/iamahless/product-filter-api/internal/dtos/filters"
	"github.com/iamahless/product-filter-api/internal/repositories"
	"github.com/iamahless/product-filter-api/internal/services/discount"
	"github.com/iamahless/product-filter-api/internal/services/discount/rules"
)

type ProductService struct {
	Repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) *ProductService {
	return &ProductService{Repo: repo}
}

func (service *ProductService) ListProducts(filters filters.ProductFilter) ([]dtos.ProductResponse, error) {
	productResponses := []dtos.ProductResponse{}
	calculator := discount.NewDiscountCalculator([]rules.DiscountRule{
		rules.CategoryDiscountRule{},
		rules.SkuDiscountRule{},
	})

	products, err := service.Repo.FindByFilters(filters)

	if err != nil {
		return nil, err
	}

	for _, product := range products {
		var discountString string
		var finalPrice int
		originalPrice := product.Price
		discountPercentage := calculator.GetBestDiscount(product)

		if discountPercentage != nil {
			finalPrice = int(math.Round(float64(originalPrice) * (1 - float64(*discountPercentage)/100)))

			fmt.Println("Discount applied:", *discountPercentage)
			fmt.Println("Final price:", finalPrice)
			discountString = fmt.Sprintf("%.0f%%", float64(*discountPercentage))
		} else {
			finalPrice = originalPrice
			discountString = ""
		}

		productResponse := dtos.ProductResponse{
			Sku:      product.Sku,
			Name:     product.Name,
			Category: product.Category,
			Price:    dtos.NewProductPriceResponse(originalPrice, finalPrice, discountString),
		}
		productResponses = append(productResponses, productResponse)
	}

	return productResponses, nil
}
