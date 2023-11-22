package handlers

import (
	services"github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
)

type CouponHandler struct{
	couponUsecase services.CouponUsecase
}
// Constructor function
func NewCouponHandler(couponUsecase services.CouponUsecase)*CouponHandler{
	return &CouponHandler{
		couponUsecase: couponUsecase,
	}
}