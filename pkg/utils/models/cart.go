package models

import "github.com/Anandhu4456/go-Ecommerce/pkg/domain"

type GetCart struct {
	ProductName   string  `json:"product_name"`
	CategoryId    int     `json:"category_id"`
	Quantity      int     `json:"quantity"`
	Total         float64 `json:"total"`
	DiscoundPrice float64 `json:"discount_price"`
}

type CheckOut struct {
	Addresses []domain.Address
	Products []GetCart
	PaymentMethods []domain.PaymentMethod
	TotalPrice float64
}

type Order struct{
	AddressId int `json:"address_id"`
	PaymentId int `json:"payment_id"`
}
