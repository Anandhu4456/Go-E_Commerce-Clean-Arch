package usecase

import (
	"mime/multipart"

	"github.com/Anandhu4456/go-Ecommerce/pkg/helper"
	"github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
)

type inventoryUsecase struct {
	invRepo interfaces.InventoryRespository
}

// constructor function
func NewInventoryUsecase(invRepo interfaces.InventoryRespository) services.InventoryUsecase {
	return &inventoryUsecase{
		invRepo: invRepo,
	}
}

func (invU *inventoryUsecase) AddInventory(inventory models.Inventory, image *multipart.FileHeader) (models.InventoryResponse, error) {
	url, err := helper.AddImageToS3(image)
	if err != nil {
		return models.InventoryResponse{}, err
	}
	inventory.Image = url

	// Send the url and save in db
	inventoryResponse, err := invU.invRepo.AddInventory(inventory, url)
	if err != nil {
		return models.InventoryResponse{}, err
	}
	return inventoryResponse, nil
}

func (invU *inventoryUsecase) UpdateImage(invID int, image *multipart.FileHeader) (models.Inventory, error) {
	url, err := helper.AddImageToS3(image)
	if err != nil {
		return models.Inventory{}, err
	}

	inventoryResponse, err := invU.invRepo.UpdateImage(invID, url)
	if err != nil {
		return models.Inventory{}, err
	}
	return inventoryResponse, nil
}
