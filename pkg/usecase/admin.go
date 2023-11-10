package usecase

import (
	"errors"

	"github.com/Anandhu4456/go-Ecommerce/pkg/helper"
	"github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
	"golang.org/x/crypto/bcrypt"
)

type adminUsecase struct {
	adminRepository interfaces.AdminRepository
}

// constructor function
func NewAdminUsecase(adRepo interfaces.AdminRepository) services.AdminUsecase {
	return &adminUsecase{
		adminRepository: adRepo,
	}
}

func (au *adminUsecase) LoginHandler(adminDetails models.AdminLogin) (models.AdminToken, error) {
	// Getting details of the admin based on the email provided

	adminCompareDetails, err := au.adminRepository.LoginHandler(adminDetails)
	if err != nil {
		return models.AdminToken{}, errors.New("admin not found")
	}
	// Compare password from database that provided by admin

	err = bcrypt.CompareHashAndPassword([]byte(adminCompareDetails.Password), []byte(adminDetails.Password))
	if err != nil {
		return models.AdminToken{}, errors.New("password doesn't match")
	}

	token, err := helper.GenerateAdminToken(adminCompareDetails)
	if err != nil {
		return models.AdminToken{}, err
	}
	return models.AdminToken{
		Username: adminCompareDetails.Username,
		Token:    token,
	}, nil
}

func (au *adminUsecase) BlockUser(id string) error {
	user, err := au.adminRepository.GetUserById(id)
	if err != nil {
		return errors.New("user not found")
	}
	if !user.Permission {
		return errors.New("already blocked")
	} else {
		user.Permission = false
	}
	err = au.adminRepository.UpdateBlockUserById(user)
	if err != nil {
		return err
	}
	return nil
}
