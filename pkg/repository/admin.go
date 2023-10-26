package repository

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/domain"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
	"gorm.io/gorm"
)

type adminRepository struct {
	DB *gorm.DB
}

func (ar *adminRepository) LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error) {
	var adminCompareDetail domain.Admin
	err := ar.DB.Raw("select * from admins where email=?", adminDetails.Email).Scan(&adminCompareDetail).Error
	if err != nil {
		return domain.Admin{}, err
	}
	return adminCompareDetail, nil
}
