package discount

import (
	"github.com/iamahless/product-filter-api/internal/models"
	"github.com/iamahless/product-filter-api/internal/services/discount/rules"
)

type DiscountCalculator struct {
	rules []rules.DiscountRule
}

func NewDiscountCalculator(rules []rules.DiscountRule) *DiscountCalculator {
	return &DiscountCalculator{
		rules: rules,
	}
}

func (calculator *DiscountCalculator) GetBestDiscount(product models.Product) *int {
	var best *int

	for _, rule := range calculator.rules {
		discount := rule.CalculateDiscount(product)
		if discount != nil {
			if best == nil || *discount > *best {
				best = discount
			}
		}
	}

	return best
}
