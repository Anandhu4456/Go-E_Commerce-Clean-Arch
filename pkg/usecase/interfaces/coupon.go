package interfaces

import "github.com/Anandhu4456/go-Ecommerce/pkg/domain"

type CouponUsecase interface {
	Addcoupon(coupon domain.Coupon) error
	MakeCouponInvalid(id int) error
	GetCoupons(page, limit int) ([]domain.Coupon, error)
}
