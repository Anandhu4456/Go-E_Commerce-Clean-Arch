package usecase

import (
	"errors"

	"github.com/Anandhu4456/go-Ecommerce/pkg/domain"
	"github.com/Anandhu4456/go-Ecommerce/pkg/helper"
	interfaces "github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
	"github.com/jinzhu/copier"
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

func (au *adminUsecase) LoginHandler(adminDetails models.AdminLogin) (domain.AdminToken, error) {
	// Getting details of the admin based on the email provided

	adminCompareDetails, err := au.adminRepository.LoginHandler(adminDetails)
	if err != nil {
		return domain.AdminToken{}, errors.New("admin not found")
	}

	hash, err := helper.PasswordHashing(adminDetails.Password)
	if err != nil {
		return domain.AdminToken{}, err
	}
	// Compare password from database that provided by admin
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(adminDetails.Password))
	if err != nil {
		return domain.AdminToken{}, err
	}

	var adminDetailsResponse models.AdminDetailsResponse
	err = copier.Copy(&adminDetailsResponse, &adminCompareDetails)
	if err != nil {
		return domain.AdminToken{}, err
	}

	token, refresh, err := helper.GenerateAdminToken(adminDetailsResponse)
	if err != nil {
		return domain.AdminToken{}, err
	}
	return domain.AdminToken{
		Admin:        adminDetailsResponse,
		Token:        token,
		RefreshToken: refresh,
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

func (au *adminUsecase) UnblockUser(id string) error {
	user, err := au.adminRepository.GetUserById(id)
	if err != nil {
		return errors.New("user not found")
	}
	if !user.Permission {
		// means that user is blocked(false)..then,
		user.Permission = true
	}
	err = au.adminRepository.UpdateBlockUserById(user)
	if err != nil {
		return errors.New("error in unblock user")
	}
	return nil
}

func (au *adminUsecase) GetUsers(page, limit int) ([]models.UserDetailsAtAdmin, error) {
	users, err := au.adminRepository.GetUsers(page, limit)
	if err != nil {
		return []models.UserDetailsAtAdmin{}, errors.New("getting users failed")
	}
	return users, nil

}
