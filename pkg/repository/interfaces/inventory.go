package interfaces

import "github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"

type InventoryRespository interface{
	AddInventory(inventory models.Inventory, url string)(models.InventoryResponse,error)
	UpdateInventory(pid int, invData models.UpdateInventory)(models.Inventory,error)
	DeleteInventory(id string)error
	CheckInventory(pid int)(bool,error)

	AddImage(product_id int,image_url string)(models.InventoryResponse,error)
	UpdateImage(invId int,url string)(models.Inventory,error)
	DeleteImage(product_id int,imageId int)error
	GetImagesFromInventoryId(product_id int,)([]models.ImagesInfo,error)

	ListProducts(page,limit int)([]models.InventoryList,error)
	ShowIndividualProducts(id string)(models.Inventory,error)
	SearchProducts(key string,page,limit int)([]models.InventoryList,error)
	GetCategoryProducts(categoryId,page,limit int)([]models.InventoryList,error)

	CheckStock(inventory_id int)(int,error)
	CheckPrice(inventory_id int)(float64,error)
}