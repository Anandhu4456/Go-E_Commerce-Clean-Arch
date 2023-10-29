package repository

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	"gorm.io/gorm"
)


type couponRepository struct{
	DB *gorm.DB
}

// constructor function
func NewCouponRepository(DB *gorm.DB)interfaces.CouponRepository{
	return &couponRepository{
		DB:DB,
	}
}