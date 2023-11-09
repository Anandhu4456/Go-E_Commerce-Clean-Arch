package interfaces

import (
	"mime/multipart"

	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
)

type InventoryUsecase interface {
	AddInventory(inventory models.Inventory, image *multipart.FileHeader) (models.InventoryResponse, error)
	UpdateInventory(invID int, invData models.UpdateInventory) (models.Inventory, error)
	UpdateImage(invID int, image *multipart.FileHeader) (models.Inventory, error)
	DeleteInventory(id string) error

	ShowIndividualProducts(s string) (models.InventoryDetails, error)
	ListProducts(page int, limit int) ([]models.InventoryList, error)
	SearchProducts(key string, page, limit int) ([]models.InventoryList, error)
	GetCategoryProducts(catID int, page, limit int) ([]models.InventoryList, error)
	AddImage(product_id int, image *multipart.FileHeader) (models.InventoryResponse, error)
	DeleteImage(product_id, image_id int) error
}
