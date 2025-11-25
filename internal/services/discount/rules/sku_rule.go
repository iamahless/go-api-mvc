package rules

import "github.com/iamahless/product-filter-api/internal/models"

type SkuDiscountRule struct{}

func (rule SkuDiscountRule) CalculateDiscount(product models.Product) *int {
	if product.Sku == "000003" {
		v := 15
		return &v
	}

	return nil
}
