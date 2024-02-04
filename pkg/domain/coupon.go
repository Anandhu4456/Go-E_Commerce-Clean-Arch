package domain

type Coupon struct {
	Coupon string `json:"coupon" gorm:"unique;not null"`
	DiscountRate int    `json:"discount_rate"`
	Valid        bool   `gorm:"default:true"`
}
