package repository

import (
	"errors"

	"github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
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
	if err := wlr.DB.Raw("SELECT id FROM wishlist WHERE user_id=?", user_id).Scan(&wishlistId).Error; err != nil {
		return 0, errors.New("wishlist id not found")
	}
	return wishlistId, nil
}

func (wlr *wishlistReposotory) GetWishlist(id int) ([]models.GetWishlist, error) {
	var getWishlist []models.GetWishlist

	query := `
	SELECT wishlist.user_id,categories.category,inventories.product_name,inventories.price
	FROM wishlist
	JOIN wishlist_items.wishlist_id=wishlist.idgetWishlist
	JOIN invenotries ON wishlist_items.inventory_id=inventories.id
	JOIN categories ON inventories.category_id=categories.id 
	WHERE wishlist.user_id

	`
	if err := wlr.DB.Raw(query, id).Scan(&getWishlist).Error; err != nil {
		return []models.GetWishlist{}, err
	}
	return getWishlist, nil
}

func (wlr *wishlistReposotory) GetProductsInWishlist(wishlistId int) ([]int, error) {
	var productsInWishlist []int

	if err := wlr.DB.Raw("SELECT inventory_id FROM wishlist_items WHERE wishlist_id", wishlistId).Scan(&productsInWishlist).Error; err != nil {
		return []int{}, err
	}
	return productsInWishlist, nil
}
