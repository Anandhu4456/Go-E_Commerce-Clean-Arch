package repository

import (
	"errors"
	"fmt"

	"github.com/Anandhu4456/go-Ecommerce/pkg/domain"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
	"gorm.io/gorm"
)

type couponRepository struct {
	DB *gorm.DB
}

// constructor function
func NewCouponRepository(DB *gorm.DB) *couponRepository {
	return &couponRepository{
		DB: DB,
	}
}

func (couprep *couponRepository) AddCoupon(coupon models.Coupon) error {
	err := couprep.DB.Exec("INSERT INTO coupons(name,discount_rate,valid)VALUES ($1,$2,$3)", coupon.DiscountRate, coupon.Valid).Error
	if err != nil {
		return err
	}
	return nil
}

func (couprep *couponRepository) MakeCouponInvalid(id int) error {
	err := couprep.DB.Exec("UPDATE coupons SET valid=false WHERE id=$1", id).Error
	if err != nil {
		return err
	}
	return nil
}

func (couprep *couponRepository) FindCouponDiscount(coupon string) int {
	var discountRate int

	err := couprep.DB.Raw("SELECT discount_rate FROM coupons WHERE name=?", coupon).Scan(&discountRate).Error
	if err != nil {
		return 0
	}
	fmt.Println("discount rate: ", discountRate)
	return discountRate
}

func (couprep *couponRepository) GetCoupons(page, limit int) ([]domain.Coupon, error) {
	var coupon []domain.Coupon
	err := couprep.DB.Raw("SELECT * FROM coupons").Scan(&coupon).Error
	if err != nil {
		return []domain.Coupon{}, err
	}
	return coupon, nil
}

func (couprep *couponRepository) ValidateCoupon(coupon string) (bool, error) {
	count := 0
	err := couprep.DB.Raw("SELECT COUNT(id)FROM coupons WHERE name=?", coupon).Scan(&count).Error
	if err != nil {
		return false, err
	}
	if count < 1 {
		return false, errors.New("not a valid coupon")
	}
	valid := true
	err = couprep.DB.Raw("SELECT valid FROM coupons WHERE name = ?", coupon).Scan(&valid).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
