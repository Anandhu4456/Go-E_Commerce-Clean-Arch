package models

type InventoryResponse struct {
	ProductID int
}

type Inventory struct {
	ID          uint    `json:"id"`
	CategoryID  int     `json:"category_id"`
	Image       string  `json:"image"`
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}

type AddToCart struct {
	UserID      int `json:"user_id"`
	InventoryID int `json:"inventory_id"`
}

type UpdateInventory struct {
	CategoryID  int     `json:"category_id"`
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}

type InventoryList struct {
	ID          uint    `json:"id"`
	CategoryID  int     `json:"category_id"`
	Image       string  `json:"image"`
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}

type ImagesInfo struct {
	ID       int    `json:"id"`
	ImageUrl string `json:"image_url"`
}

type InventoryDetails struct {
	Inventory        Inventory
	AdditionalImages []ImagesInfo
}
