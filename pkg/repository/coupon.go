package repository

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/domain"
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

func (couprep *couponRepository)AddCoupon(coupon domain.Coupon)error{
	err:=couprep.DB.Exec("INSERT INTO coupons(name,discount_rate,valid)VALUES ($1,$2,$3)",coupon.Name,coupon.DiscountRate,coupon.Valid).Error
	if err!=nil{
		return err
	}
	return nil
}

func (couprep *couponRepository)MakeCouponInvalid(id int)error{
	err:=couprep.DB.Exec("UPDATE coupons SET valid=false WHERE id=$1",id).Error
	if err!=nil{
		return err
	}
	return nil
}