package handlers

import (
	"net/http"
	// "strconv"

	
	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	CategoryUsecase services.CategoryUsecase
}

// Constructor function

func NewCategoryHandler(categoryUsecase services.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{
		CategoryUsecase: categoryUsecase,
	}
}

// @Summary		Add Category
// @Description	Admin can add new categories for products
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			category	query	string	true	"category"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/category/add [post]
func (catH *CategoryHandler) AddCategory(c *gin.Context) {
	// cat := c.Query("category")
	var cat string
	if err:=c.BindJSON(&cat);err!=nil{
		errRes:=response.ClientResponse(http.StatusBadRequest,"fileds provided are in wrong format",nil,err.Error())
		c.JSON(http.StatusBadRequest,errRes)
		return
	}
	categoryRes, err := catH.CategoryUsecase.AddCategory(cat)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Couldn't add category", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Category added successfully", categoryRes, nil)
	c.JSON(http.StatusOK, successRes)
	
}

// @Summary		Update Category
// @Description	Admin can update name of a category into new name
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			set_new_name	body	models.SetNewName	true	"set new name"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/category/update [patch]
func (catH *CategoryHandler) UpdateCategory(c *gin.Context) {
	var updateCategory models.SetNewName

	if err := c.BindJSON(&updateCategory); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "field provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	updateRes, err := catH.CategoryUsecase.UpdateCategory(updateCategory.Current, updateCategory.New)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "update category failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "category updated successfully", updateRes, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Delete Category
// @Description	Admin can delete a category
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			id	query	string	true	"id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/category/delete [delete]
func (catH *CategoryHandler) DeleteCategory(c *gin.Context) {
	// Find category id
	categoryId := c.Query("id")
	if err := catH.CategoryUsecase.DeleteCategory(categoryId); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "category deleted", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		List Categories
// @Description	Admin can view the list of  Categories
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/category [get]
func (catH *CategoryHandler) Categories(c *gin.Context) {

	categories, err := catH.CategoryUsecase.GetCategories()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully retrieved categories", categories, nil)
	c.JSON(http.StatusOK, successRes)
}
