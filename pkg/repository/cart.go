package repository

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type cartRepository struct{
	DB *gorm.DB
}

// constructor function
func NewCartRepository (DB *gorm.DB)interfaces.CartRepository{
	return &cartRepository{
		DB:DB,
	}
}

