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

// @Summary		Add Inventory
// @Description	Admin can add new  products
// @Tags			Admin
// @Accept			multipart/form-data
// @Produce		    json
// @Param			category_id		formData	string	true	"category_id"
// @Param			product_name	formData	string	true	"product_name"
// @Param			description		formData	string	true	"description"
// @Param			price	formData	string	true	"price"
// @Param			stock		formData	string	true	"stock"
// @Param           image      formData     file   true   "image"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/inventories/add [post]
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

// @Summary		Add image to an Inventory
// @Description	Admin can add new image to product
// @Tags			Admin
// @Accept			multipart/form-data
// @Produce		    json
// @Param			product_id	formData	string	true	"product_id"
// @Param           image      formData     file   true   "image"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/inventories/add-image [post]
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

// @Summary		Delete Inventory image
// @Description	Admin can delete a product image
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			product_id	query	string	true	"product_id"
// @Param			image_id	query	string	true	"image_id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/inventories/delete-image [delete]
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

// @Summary		Update inventory
// @Description	Admin can update inventory details
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			id	query	string	true	"id"	
// @Param			updateinventory	body	models.UpdateInventory	true	"Update Inventory"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/inventories/update [patch]
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

// @Summary		Update image
// @Description	Admin can update image of the inventory
// @Tags			Admin
// @Accept			multipart/form-data
// @Produce		    json
// @Param			id	query	string	true	"id"	
// @Param           image      formData     file   true   "image"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/inventories/update-image [patch]
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

// @Summary		Delete Inventory
// @Description	Admin can delete a product
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			id	query	string	true	"id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/inventories/delete [delete]
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

// @Summary		Show Product Details
// @Description	client can view the details of the product
// @Tags			Products
// @Accept			json
// @Produce		    json
// @Param			inventoryID	query	string	true	"Inventory ID"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/products/details [get]
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

// @Summary		List Products
// @Description	client can view the list of available products
// @Tags			Products
// @Accept			json
// @Produce		    json
// @Param			page	query  string 	true	"page"
// @Param			limit	query  string 	true	"limit"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/products [get]
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

// @Summary		Search Products
// @Description	client can search with a key and get the list of  products similar to that key
// @Tags			Products
// @Accept			json
// @Produce		    json
// @Param			page	query  string 	true	"page"
// @Param			limit	query  string 	true	"limit"
// @Param			searchkey 	query  string 	true	"searchkey"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/products/search [get]
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

// @Summary		List Products
// @Description	client can view the list of available products
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			page	query  string 	true	"page"
// @Param			limit	query  string 	true	"limit"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/products [get]
func (invH *InventoryHandler) AdminListProdutcs(c *gin.Context) {
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

// @Summary		filter Products by category
// @Description	client can filter with a category and get the list of  products in the category
// @Tags			Products
// @Accept			json
// @Produce		    json
// @Param			page	query  string 	true	"page"
// @Param			limit	query  string 	true	"limit"
// @Param			catID 	query  string 	true	"category ID"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/products/category [get]
func (invH *InventoryHandler) GetCategoryProducts(c *gin.Context) {
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
	catIdStr := c.Query("cat_id")
	catId, err := strconv.Atoi(catIdStr)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "category id is in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	result, err := invH.inventoryUsecase.GetCategoryProducts(catId, page, limit)
	if err != nil {
		errRes := response.ClientResponse(http.StatusOK, "couldn't get category products", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully get category products", result, nil)
	c.JSON(http.StatusOK, successRes)
}
