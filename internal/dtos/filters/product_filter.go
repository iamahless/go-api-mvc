package filters

type ProductFilter struct {
	Category      string `form:"category"`
	PriceLessThan int    `form:"priceLessThan"`
	Limit         int    `form:"limit"`
	Page          int    `form:"page"`
}

func NewProductFilter() ProductFilter {
	return ProductFilter{
		Limit: 20,
		Page:  1,
	}
}
