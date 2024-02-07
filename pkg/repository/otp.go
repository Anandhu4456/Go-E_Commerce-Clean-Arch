package repository

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
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

func (otr *otpRepository) UserDetailsUsingPhone(phone string) (models.UserDetailsResponse, error) {
	var userDetails models.UserDetailsResponse

	err := otr.DB.Raw("SELECT * FROM users WHERE phone=?", phone).Scan(&userDetails).Error
	if err != nil {
		return models.UserDetailsResponse{}, err
	}
	return userDetails, nil
}
