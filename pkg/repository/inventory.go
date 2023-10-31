package repository

import (
	"errors"
	"strconv"

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

func (ir *inventoryRespository) DeleteInventory(inventoryId string) error {
	id, err := strconv.Atoi(inventoryId)
	if err != nil {
		return errors.New("string to int conversion failed")
	}
	result := ir.DB.Exec("DELETE FROM inventories WHERE id = ? ", id)

	if result.RowsAffected < 1 {
		return errors.New("no records exists with this id")
	}
	return nil
}

func (ir *inventoryRespository) ShowIndividualProducts(id string) (models.Inventory, error) {
	pid, err := strconv.Atoi(id)
	if err != nil {
		return models.Inventory{}, errors.New("string to int conversion failed")
	}
	var product models.Inventory

	err = ir.DB.Raw(`
		SELECT * FROM inventories
		WHERE Inventories.id = ?
	`, pid).Scan(&product).Error
	if err != nil {
		return models.Inventory{}, errors.New("error occured while showing individual product")
	}
	return product, err
}

func (ir *inventoryRespository) ListProducts(page, limit int) ([]models.InventoryList, error) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	var productDetails []models.InventoryList

	err := ir.DB.Raw("SELECT inventories.id,inventories.product_name,inventories.description,inventories.stock,inventories.price,inventories.image,categories.category AS category FROM inventories JOIN categories ON inventories.category_id = categories.id LIMIT ? OFFSET ?", limit, offset).Scan(&productDetails).Error
	if err != nil {
		return []models.InventoryList{}, err
	}
	return productDetails, nil
}
func (ir *inventoryRespository) CheckStock(inventory_id int) (int, error) {
	var stock int
	err := ir.DB.Raw("SELECT stock FROM inventories WHERE id = ? ", inventory_id).Scan(&stock).Error
	if err != nil {
		return 0, err
	}
	return stock, nil
}

func (ir *inventoryRespository) CheckPrice(inventory_id int) (float64, error) {
	var price float64
	err := ir.DB.Raw("SELECT price FROM inventories WHERE id = ?", inventory_id).Scan(&price).Error
	if err != nil {
		return 0, err
	}
	return price, err
}

func (ir *inventoryRespository) SearchProducts(key string, page, limit int) ([]models.InventoryList, error) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	var productSearchResult []models.InventoryList

	query := `
	
	SELECT 
		inventories.id,
		inventories.product_name,
		inventories.description,
		inventories.stock,
		inventories.price,
		inventories.image,
		categories.category AS category
	FROM
		inventories
	JOIN
		categories
	ON
		inventories.category_id = categories.id
	WHERE
		product_name ILIKE '%' || ? || '%'
	OR description ILIKE '%' || ? || '%'
	LIMIT ? OFFSET ?
	`
	err := ir.DB.Raw(query, key, limit, offset).Scan(&productSearchResult).Error
	if err != nil {
		return []models.InventoryList{}, err
	}
	return productSearchResult, nil
}

func (ir *inventoryRespository) GetCategoryProducts(categoryId, page, limit int) ([]models.InventoryList, error) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	var categoryProd []models.InventoryList

	query := `
	
	SELECT 
		inventories.id,
		inventories.product_name,
		inventories.description,
		inventories.image,
		inventories.stock,
		inventories.price
		categories.categorie AS category
	FROM
		inventories 
	JOIN
		categories
	ON
		inventories.category_id = categories.id
	WHERE 
		inventories.category_id = ?
	LIMIT ? OFFSET ?
	`

	err := ir.DB.Raw(query, categoryId, limit, offset).Scan(&categoryProd).Error
	if err != nil {
		return []models.InventoryList{}, err
	}
	return categoryProd, nil
}

func (ir *inventoryRespository) AddImage(product_id int, image_url string) (models.InventoryResponse, error) {
	var addImageResponse models.InventoryResponse

	query := `
	
	INSERT INTO 
		images (inventory_id,image_url)
	VALUES (?,?)
	RETURNING 
		inventory_id
	`
	err := ir.DB.Raw(query, product_id, image_url).Scan(&addImageResponse).Error
	if err != nil {
		return models.InventoryResponse{}, errors.New("adding image failed")
	}
	return addImageResponse, nil
}

func (ir *inventoryRespository) DeleteImage(product_id int, imageId int) error {
	result := ir.DB.Exec("DELETE FROM images WHERE id= ?", imageId)

	if result.RowsAffected < 1 {
		return errors.New("no image exists with this id")
	}
	return nil
}

func (ir *inventoryRespository) GetImagesFromInventoryId(product_id int) ([]models.ImagesInfo, error) {
	var imagesFromInvId []models.ImagesInfo
	err := ir.DB.Raw("SELECT id,image_url FROM images WHERE inventory_id = ?", product_id).Scan(&imagesFromInvId).Error
	if err != nil {
		return []models.ImagesInfo{}, err
	}
	return imagesFromInvId, nil
}
