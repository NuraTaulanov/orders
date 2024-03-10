package models

type NewOrderReq struct {
	CustomerName string    `json:"customer_name"`
	Products     []Product `json:"products"`
}

type OrderResponse struct {
	OrderId      int64  `json:"order_id"`
	ProductId    int64  `json:"product_id"`
	ProductEan   string `json:"product_ean"`
	CustomerName string `json:"customer_name"`
}
