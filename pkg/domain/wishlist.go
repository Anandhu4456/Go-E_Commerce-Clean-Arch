package domain

// Wishlist of a user

type WishList struct {
	ID     uint `json:"id" gorm:"primarykey"`
	UserID uint `json:"user_id" gorm:"not null"`
	User   User `json:"-" gorm:"foreignkey:UserID"`
}

// Products in the wishlist of user

type WishlistItems struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	WishlistID  uint      `json:"cart_id" gorm:"not null"`
	Wishlist    WishList  `json:"-" gorm:"foreignkey:WishlistID"`
	InventoryID uint      `json:"inventory_id" gorm:"not null"`
	Inventory   Inventory `json:"-" gorm:"foreignkey:InventoryID"`
}
