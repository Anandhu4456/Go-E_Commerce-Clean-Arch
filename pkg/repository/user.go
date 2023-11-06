package repository

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	"gorm.io/gorm"
)


type userRepository struct{
	DB *gorm.DB
}

// constructor funciton

func NewUserRepository(DB *gorm.DB)interfaces.UserRepository{
	return &userRepository{
		DB:DB,
	}
}

func (ur *userRepository)CheckUserAvailability(email string) bool{
	var userCount int

	err:=ur.DB.Raw("SELECT COUNT(*) FROM users WHERE email=?",email).Scan(&userCount).Error
	if err!=nil{
		return false
	}
	// if count greater than 0, user already exist
	return userCount >0
}