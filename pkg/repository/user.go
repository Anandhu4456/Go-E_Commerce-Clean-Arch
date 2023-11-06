package repository

import (
	"errors"

	"github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

// constructor funciton

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userRepository{
		DB: DB,
	}
}

func (ur *userRepository) CheckUserAvailability(email string) bool {
	var userCount int

	err := ur.DB.Raw("SELECT COUNT(*) FROM users WHERE email=?", email).Scan(&userCount).Error
	if err != nil {
		return false
	}
	// if count greater than 0, user already exist
	return userCount > 0
}

func (ur *userRepository) UserBlockStatus(email string) (bool, error) {
	var permission bool
	err := ur.DB.Raw("SELECT permission FROM users WHERE email=?", email).Scan(&permission).Error
	if err != nil {
		return false, err
	}
	return permission, nil
}

func (ur *userRepository) FindUserByEmail(user models.UserLogin) (models.UserResponse, error) {
	var userResponse models.UserResponse
	err := ur.DB.Raw("SELECT * FROM users WHERE email=? AND permission=true", user.Email).Scan(&userResponse).Error
	if err != nil {
		return models.UserResponse{}, errors.New("no user found")
	}
	return userResponse, nil
}
