package custom_model

type UpdateProduct struct {
	ID          int    `json:"id"`
	ProductCode string `json:"productCode"`
	ProductName string `json:"productName"`
	Price       int    `json:"price"`
}
