package domain

// cart of a user
type Cart struct {
	ID     uint `json:"id" gorm:"primarykey"`
	UserID uint `json:"user_id" gorm:"not null"`
	User   User `json:"user" gorm:"foreignkey:UserID"`
}

// products in the cart of user
type LineItems struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	CartID      uint      `json:"cart_id" gorm:"not null"`
	Cart        Cart      `json:"-" gorm:"foreignkey:CartID"`
	InventoryID uint      `json:"inventory_id" gorm:"not null"`
	Inventory   Inventory `json:"-" gorm:"foreignkey:InventoryID"`
	Quantity    int       `json:"quantity" gorm:"default:1"`
}
