package rules

import "github.com/iamahless/product-filter-api/internal/models"

type DiscountRule interface {
	CalculateDiscount(product models.Product) *int
}
