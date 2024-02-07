package models

type OrderPaymentDetails struct {
	UserID     int     `json:"user_id"`
	Username   string  `json:"username"`
	RazorID    int     `json:"razor_id"`
	OrderID    int     `json:"order_id"`
	FinalPrice float64 `json:"final_price"`
}

type Order struct {
	UserID          int `json:"user_id"`
	AddressID       int `json:"address_id"`
	PaymentMethodID int `json:"payment_id"`
	CouponID        int `json:"coupon_id"`
}

type InvoiceData struct {
	Title       string
	Quantity    int
	Price       int
	TotalAmount int
}

type Invoice struct {
	Name         string
	Address      string
	InvoiceItems []*InvoiceData
}
