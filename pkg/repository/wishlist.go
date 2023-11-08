package repository

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type wishlistReposotory struct{
	DB *gorm.DB
}

// constructor function 

func NewWishlistRepository(DB *gorm.DB)interfaces.WishlistRepository{
	return &wishlistReposotory{
		DB:DB,
	}
}