package domain

import "gorm.io/gorm"

type Coupon struct {
	gorm.Model
	Coupon       string `json:"coupon" gorm:"unique;not null"`
	DiscountRate int    `json:"discount_rate"`
	Valid        bool   `gorm:"default:true"`
}
