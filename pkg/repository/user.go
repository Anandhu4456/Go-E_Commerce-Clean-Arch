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