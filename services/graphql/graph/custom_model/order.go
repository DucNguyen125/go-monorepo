package custom_model

type UpdateOrder struct {
	ID          int    `json:"id"`
	OrderCode   string `json:"orderCode"`
	OrderType   string `json:"orderType"`
	Products    string `json:"products"`
	OrderStatus string `json:"orderStatus"`
	Quantity    int    `json:"quantity"`
	TotalPrice  int    `json:"totalPrice"`
}
