package usecase

import (
	"errors"

	"github.com/Anandhu4456/go-Ecommerce/pkg/domain"
	interfaces "github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
)

type couponUsecase struct {
	couponRepo interfaces.CouponRepository
}

// constructor function

func NewCouponUsecase(couponRepo interfaces.CouponRepository) services.CouponUsecase {
	return &couponUsecase{
		couponRepo: couponRepo,
	}
}

func (coupU *couponUsecase) Addcoupon(coupon domain.Coupon) error {
	if err := coupU.couponRepo.AddCoupon(coupon); err != nil {
		return errors.New("coupon adding failed")
	}
	return nil
}
