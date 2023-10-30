package repository

import (
	"errors"

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

func (ir *inventoryRespository) UpdateImage(inventId int, url string) (models.Inventory, error) {

	// check db connecton
	if ir.DB == nil {
		return models.Inventory{}, errors.New("database connection failed while updating image")
	}

	// updating image
	err := ir.DB.Exec("UPDATE inventories SET image=? WHERE id= ? ", url, inventId).Error

	if err != nil {
		return models.Inventory{}, err
	}
	// Retrieve the update
	var updatedImageInventory models.Inventory
	err = ir.DB.Raw("SELECT * FROM inventories WHERE id = ?", inventId).Scan(&updatedImageInventory).Error
	if err != nil {
		return models.Inventory{}, err
	}
	return updatedImageInventory, nil
}

func (ir *inventoryRespository) CheckInventory(pid int) (bool, error) {
	var check int
	err := ir.DB.Raw("SELECT COUNT(*)FROM inventories WHERE id = ?", pid).Scan(&check).Error
	if err != nil {
		return false, err
	}
	if check == 0 {
		return false, err
	}
	return true, nil
}

func (ir *inventoryRespository) UpdateInventory(pid int, invData models.UpdateInventory) (models.Inventory, error) {
	if ir.DB == nil {
		return models.Inventory{}, errors.New("database connection failed while update inventory")
	}

	if invData.CategoryID != 0 {
		if err := ir.DB.Exec("UPDATE inventories SET categorie_id WHERE id=?", invData.CategoryID, pid).Error; err != nil {
			return models.Inventory{}, err
		}
	}

	if invData.ProductName != "" && invData.ProductName != "string" {
		if err := ir.DB.Exec("UPDATE inventories SET product_name=? WHERE id= ?", invData.ProductName, pid).Error; err != nil {
			return models.Inventory{}, err
		}
	}

	if invData.Description != "" && invData.Description != "string" {
		if err := ir.DB.Exec("UPDATE inventories SET description=? WHERE id=?", invData.Description, pid).Error; err != nil {
			return models.Inventory{}, err
		}
	}

	if invData.Stock != 0 {
		if err := ir.DB.Exec("UPDATE inventories SET stock= ? WHERE id=?", invData.Stock, pid).Error; err != nil {
			return models.Inventory{}, err
		}
	}

	if invData.Price != 0 {
		if err := ir.DB.Exec("UPDATE inventories SET price=? WHERE id=?", invData.Price, pid).Error; err != nil {
			return models.Inventory{}, err
		}
	}

	// retrieve the updates
	var updatedInventory models.Inventory
	err := ir.DB.Raw("SELECT * FROM inventories WHERE Id = ?", pid).Scan(&updatedInventory).Error
	if err != nil {
		return models.Inventory{}, err
	}
	return updatedInventory, nil
}
