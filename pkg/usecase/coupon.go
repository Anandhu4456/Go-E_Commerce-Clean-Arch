package usecase

import (
	interfaces "github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
)

type couponUsecase struct {
	couponRepo interfaces.CouponRepository
}

// constructor function

func NewCouponUsecase(couponRepo interfaces.CouponRepository) services.CouponUsecase{
	return &couponUsecase{
		couponRepo: couponRepo,
	}
}
