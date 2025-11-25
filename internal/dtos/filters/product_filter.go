package filters

type ProductFilter struct {
	Category      string `json:"category"`
	PriceLessThan int    `json:"price_less_than"`
	Limit         int    `json:"limit"`
}
