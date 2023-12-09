package usecase

import (
	"errors"
	"mime/multipart"
	"strconv"

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

func (invU *inventoryUsecase) UpdateInventory(invID int, invData models.UpdateInventory) (models.Inventory, error) {
	result, err := invU.invRepo.CheckInventory(invID)
	if err != nil {
		return models.Inventory{}, err
	}
	if !result {
		return models.Inventory{}, errors.New("there is no inventory as you mentioned")
	}
	newinventory, err := invU.invRepo.UpdateInventory(invID, invData)
	if err != nil {
		return models.Inventory{}, err
	}
	return newinventory, nil
}

func (invU *inventoryUsecase) DeleteInventory(id string) error {
	if err := invU.invRepo.DeleteInventory(id); err != nil {
		return err
	}
	return nil
}

func (invU *inventoryUsecase) ShowIndividualProducts(id string) (models.InventoryDetails, error) {
	product, err := invU.invRepo.ShowIndividualProducts(id)
	if err != nil {
		return models.InventoryDetails{}, err
	}
	productId, err := strconv.Atoi(id)
	if err != nil {
		return models.InventoryDetails{}, err
	}
	var AdditionalImage []models.ImagesInfo
	AdditionalImage, err = invU.invRepo.GetImagesFromInventoryId(productId)
	if err != nil {
		return models.InventoryDetails{}, err
	}
	InvDetails := models.InventoryDetails{Inventory: product, AdditionalImages: AdditionalImage}
	return InvDetails, nil
}

func (invU *inventoryUsecase) ListProducts(page int, limit int) ([]models.InventoryList, error) {
	productDetails, err := invU.invRepo.ListProducts(page, limit)

	if err != nil {
		return []models.InventoryList{}, err
	}
	return productDetails, nil
}

func (invU *inventoryUsecase) SearchProducts(key string, page, limit int) ([]models.InventoryList, error) {
	productsDetails, err := invU.invRepo.SearchProducts(key, page, limit)
	if err != nil {
		return []models.InventoryList{}, err
	}
	return productsDetails, nil
}

func (invU *inventoryUsecase) GetCategoryProducts(catID int, page, limit int) ([]models.InventoryList, error) {
	productsDetails, err := invU.invRepo.GetCategoryProducts(catID, page, limit)
	if err != nil {
		return []models.InventoryList{}, err
	}
	return productsDetails, nil
}

func (invU *inventoryUsecase) AddImage(product_id int, image *multipart.FileHeader) (models.InventoryResponse, error) {
	// adding the image to Aws s3 bucket
	imageUrl, err := helper.AddImageToS3(image)
	if err!=nil{
		return models.InventoryResponse{},err
	}

	inventoryResponse, err := invU.invRepo.AddImage(product_id, imageUrl)
	if err != nil {
		return models.InventoryResponse{}, err
	}
	return inventoryResponse, nil
}

func (invU *inventoryUsecase) DeleteImage(product_id, image_id int) error {
	if err := invU.invRepo.DeleteImage(product_id, image_id); err != nil {
		return errors.New("image not deleted")
	}
	return nil
}
