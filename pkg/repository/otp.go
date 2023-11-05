package repository

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type otpRepository struct {
	DB *gorm.DB
}

// constructor function

func NewOtpRepository(DB *gorm.DB) interfaces.OtpRepository {
	return &otpRepository{
		DB: DB,
	}
}

func (otr *otpRepository) FindUserByMobileNumber(phone string) bool {
	var count int

	err := otr.DB.Raw("SELECT COUNT(*) FROM users WHERE phone=?", phone).Scan(&count).Error
	if err != nil {
		return false
	}
	return count > 0
}
