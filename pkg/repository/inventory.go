package repository

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
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
