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

func (wlr *wishlistReposotory) FindProductNames(inventory_id int) (string, error) {
	var productName string

	if err := wlr.DB.Raw("SELECT product_name FROM inventories WHERE inventory_id=?", inventory_id).Scan(&productName).Error; err != nil {
		return "", errors.New("product name not found")
	}
	return productName, nil
}

func (wlr *wishlistReposotory) FindPrice(inventory_id int) (float64, error) {
	var price float64
	if err := wlr.DB.Raw("SELECT price FROM inventories WHERE inventory_id=?", inventory_id).Scan(&price).Error; err != nil {
		return 0, errors.New("price not found")
	}
	return price, nil
}

func (wlr *wishlistReposotory) FindCategory(inventory_id int) (string, error) {
	var category string

	query := `
	
	SELECT categories.category FROM invenotries
	JOIN categories ON inventories.category_id=category.id
	WHERE inventory_id=?

	`
	if err := wlr.DB.Raw(query, inventory_id).Scan(&category).Error; err != nil {
		return "", errors.New("category not found")
	}
	return category, nil
}
