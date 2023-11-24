package handlers

import (
	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
)

type InventoryHandler struct {
	inventoryUsecase services.InventoryUsecase
}

// Constructor funciton

func NewInventoryHandler(inventoryUsecase services.InventoryUsecase) *InventoryHandler {
	return &InventoryHandler{
		inventoryUsecase: inventoryUsecase,
	}
}
