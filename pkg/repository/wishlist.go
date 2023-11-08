package repository

import (
	"errors"

	"github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type wishlistReposotory struct {
	DB *gorm.DB
}

// constructor function

func NewWishlistRepository(DB *gorm.DB) interfaces.WishlistRepository {
	return &wishlistReposotory{
		DB: DB,
	}
}

func (wlr *wishlistReposotory) GetWishlistId(user_id int) (int, error) {
	var wishlistId int
	if err := wlr.DB.Raw("SELECT id FROM wishlists WHERE user_id=?", user_id).Scan(&wishlistId).Error; err != nil {
		return 0, errors.New("wishlist id not found")
	}
	return wishlistId, nil
}
