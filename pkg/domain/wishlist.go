package domain

// Wishlist of a user

type Wishlist struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	User        User      `json:"-" gorm:"foreignkey:UserID"`
	InventoryID uint      `json:"inventory_id" gorm:"not null"`
	Inventories Inventory `json:"-" gorm:"foreignkey:InventoryID"`
	IsDeleted   bool      `json:"is_deleted" gorm:"default:false"`
}
