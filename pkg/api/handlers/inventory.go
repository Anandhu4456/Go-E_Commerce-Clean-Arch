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

	successRes := response.ClientResponse(http.StatusOK, "deleted image successfully",nil, nil)
	c.JSON(http.StatusOK, successRes)
}
