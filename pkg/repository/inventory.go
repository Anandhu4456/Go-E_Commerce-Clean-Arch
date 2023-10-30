package repository

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
	"gorm.io/gorm"
)

type inventoryRespository struct {
	DB *gorm.DB
}

// Constructor function

func NewInventoryRepository(DB *gorm.DB) interfaces.InventoryRespository {
	return &inventoryRespository{
		DB: DB,
	}
}

func (ir *inventoryRespository) AddInventory(inventory models.Inventory, url string) (models.InventoryResponse, error) {
	var inventoryResp models.InventoryResponse

	query := `INSERT INTO inventories (categori_id,product_name,description,stock,price,image)
	VALUES(?,?,?,?,?,?)RETURNING id`

	err := ir.DB.Raw(query, inventory.CategoryID, inventory.ProductName, inventory.Description, inventory.Stock, inventory.Price, inventory.Image, url).Scan(&inventoryResp.ProductID).Error
	if err != nil {
		return models.InventoryResponse{}, err
	}
	return models.InventoryResponse{}, nil
}
