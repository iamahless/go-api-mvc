package dtos

type ProductResponse struct {
	Sku      string               `json:"sku"`
	Name     string               `json:"name"`
	Category string               `json:"category"`
	Price    ProductPriceResponse `json:"price"`
}

type ProductPriceResponse struct {
	Original           int     `json:"original"`
	Final              int     `json:"final"`
	DiscountPercentage *string `json:"discount_percentage"`
	Currency           string  `json:"currency"`
}

func NewProductPriceResponse(original, final int, discount string) ProductPriceResponse {
	var discountPtr *string

	if discount != "" {
		discountPtr = &discount
	}

	return ProductPriceResponse{
		Original:           original,
		Final:              final,
		DiscountPercentage: discountPtr,
		Currency:           "EUR",
	}
}
