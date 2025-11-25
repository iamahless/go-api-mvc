package filters

type ProductFilter struct {
	Category      string `json:"category" form:"category"`
	PriceLessThan int    `json:"price_less_than" form:"price_less_than"`
	Limit         int    `json:"limit" form:"limit"`
	Page          int    `json:"page" form:"page"`
}

func NewProductFilter() ProductFilter {
	return ProductFilter{
		Limit: 20,
		Page:  1,
	}
}
