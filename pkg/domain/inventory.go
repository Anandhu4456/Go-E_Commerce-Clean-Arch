package domain

// products in the website
type Inventory struct {
	ID uint `json:"id" gorm:"unique;not null"`
	CategoryID int `json:"category_id"`
	Category Category `json:"-" gorm:"foreignkey:CategoryID;constraint:OnDelete:CASCADE"`
	ProductName string `json:"product_name"`
	Description string `json:"description"`
	Image string `json:"image"`
	Stock int `json:"stock"`
	Price float64 `json:"price"`
}

// category of product
type Category struct{
	ID int `json:"id" gorm:"primarykey;not null"`
	Category string `json:"category"`
}

type Image struct{
	ID int`json:"id" gorm:"primarykey;not null"`
	InventoryID int `json:"inventory_id"`
	Inventory Inventory `json:"-" gorm:"foreignkey:InventoryID"`
	ImageUrl string `json:"image_url"`
}