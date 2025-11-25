package rules

import "github.com/iamahless/product-filter-api/internal/models"

type CategoryDiscountRule struct{}

func (rule CategoryDiscountRule) CalculateDiscount(product models.Product) *int {
	if product.Category == "boots" {
		v := 30
		return &v
	}

	return nil
}
