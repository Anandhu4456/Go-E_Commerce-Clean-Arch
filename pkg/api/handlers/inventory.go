package handlers

import (
	"net/http"
	"strconv"

	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/response"
	"github.com/gin-gonic/gin"
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

func (invH *InventoryHandler) AddInventory(c *gin.Context) {
	var inventory models.Inventory

	categoryId, err := strconv.Atoi(c.Request.FormValue("category_id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "form file error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	productName := c.Request.FormValue("product_name")
	description := c.Request.FormValue("description")

	p, err := strconv.Atoi(c.Request.FormValue("price"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "form file error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	price := float64(p)
	stock, err := strconv.Atoi(c.Request.FormValue("stock"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "form file error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	image, err := c.FormFile("image")
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "retrieving image from form failedd", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	inventory.CategoryID = categoryId
	inventory.Description = description
	inventory.ProductName = productName
	inventory.Price = price
	inventory.Stock = stock
	// inventory.Image = image

	InventoryResp, err := invH.inventoryUsecase.AddInventory(inventory, image)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't add inventory", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully added inventory", InventoryResp, nil)
	c.JSON(http.StatusOK, successRes)
}

func (invH *InventoryHandler) AddImage(c *gin.Context) {
	productId, err := strconv.Atoi(c.Request.FormValue("product_id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "form file error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	image, err := c.FormFile("image")
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "retrieving image from form error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	InventoryRes, err := invH.inventoryUsecase.AddImage(productId, image)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "adding image failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "image added inventory successfully", InventoryRes, nil)
	c.JSON(http.StatusOK, successRes)
}

func (invH *InventoryHandler) DeleteImage(c *gin.Context) {
	productId, err := strconv.Atoi(c.Query("product_id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "inv id not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	imageId, err := strconv.Atoi(c.Query("image_id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "image id not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	err = invH.inventoryUsecase.DeleteImage(productId, imageId)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't remove the image", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "deleted image successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (invH *InventoryHandler) UpdateInventory(c *gin.Context) {
	invIdStr := c.Query("id")
	invId, err := strconv.Atoi(invIdStr)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "id is not valid", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	var invData models.UpdateInventory

	if err := c.BindJSON(&invData); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	invRes, err := invH.inventoryUsecase.UpdateInventory(invId, invData)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "update inventory failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "updated inventory successfully", invRes, nil)
	c.JSON(http.StatusOK, successRes)

}

func (invH *InventoryHandler) UpdateImage(c *gin.Context) {
	invIdStr := c.Query("inventory_id")
	invId, err := strconv.Atoi(invIdStr)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "id is not valid", nil, err.Error())
		c.JSON(http.StatusOK, errRes)
		return
	}
	image, err := c.FormFile("image")
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "retrieve image from form failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	invRes, err := invH.inventoryUsecase.UpdateImage(invId, image)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "update image failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "updated image successfully", invRes, nil)
	c.JSON(http.StatusOK, successRes)
}

func (invH *InventoryHandler) DeleteInventory(c *gin.Context) {
	invId := c.Query("id")
	if err := invH.inventoryUsecase.DeleteInventory(invId); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "delete inventory failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "deleted inventory successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (invH *InventoryHandler) ShowIndividualProducts(c *gin.Context) {
	id := c.Query("inventory_id")
	products, err := invH.inventoryUsecase.ShowIndividualProducts(id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "path variables in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "products details retrieved successfully", products, nil)
	c.JSON(http.StatusOK, successRes)
}

func (invH *InventoryHandler) ListProdutcs(c *gin.Context) {
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "limit number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	products, err := invH.inventoryUsecase.ListProducts(page, limit)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully got all records", products, nil)
	c.JSON(http.StatusOK, successRes)
}

func (invH *InventoryHandler) SearchProducts(c *gin.Context) {
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "limit number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	searchKey := c.Query("searchkey")

	result, err := invH.inventoryUsecase.SearchProducts(searchKey, page, limit)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully searched products", result, nil)
	c.JSON(http.StatusOK, successRes)
}
