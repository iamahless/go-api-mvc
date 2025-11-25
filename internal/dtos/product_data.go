package dtos

type ProductResponse struct {
	Sku      string
	Name     string
	Category string
	price    ProductPriceResponse
}

type ProductPriceResponse struct {
	Original           int
	Final              int
	DiscountPercentage string
	Currency           string
}
