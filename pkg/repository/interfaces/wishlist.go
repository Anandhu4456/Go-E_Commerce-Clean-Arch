package interfaces

import (
	
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
)


type WishlistRepository interface{
	GetWishlist(id int)([]models.GetWishlist,error)
	GetWishlistId(user_id int)(int, error)
	CreateNewWishlist(user_id int)(int,error)
	AddWishlistItem(wishlistId,inventoryId int)error
	GetProductsInWishlist(wishlistId int)([]int,error)
	FindProductNames(inventory_id int)(string,error)
	FindPrice(inventory_id int)(float64,error)
	FindCategory(inventory_id int)(string,error)
	RemoveFromWishlist(wishlistId,inventoryId int)error
}