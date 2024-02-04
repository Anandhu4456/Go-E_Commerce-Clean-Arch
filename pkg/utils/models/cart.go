package models



type GetCart struct {
	Id            int     `json:"product_id"`
	ProductName   string  `json:"product_name"`
	CategoryId    int     `json:"category_id"`
	Quantity      int     `json:"quantity"`
	Total         float64 `json:"total"`
	DiscoundPrice float64 `json:"discount_price"`
}



type Order struct {
	AddressId int `json:"address_id"`
	PaymentId int `json:"payment_id"`
}
