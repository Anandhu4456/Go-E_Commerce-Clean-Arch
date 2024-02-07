package domain

import "gorm.io/gorm"

type Coupon struct {
	gorm.Model
	Coupon       string `json:"coupon" gorm:"unique"`
	DiscountRate int    `json:"discount_rate" `
	Valid        bool   `json:"valid" gorm:"default:true"`
}
