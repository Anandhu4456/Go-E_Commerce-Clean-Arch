package handlers

import (
	"net/http"

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

func (catH *CategoryHandler) AddCategory(c *gin.Context) {
	cat := c.Query("category")
	categoryRes, err := catH.CategoryUsecase.AddCategory(cat)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Couldn't add category", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Category added successfully", categoryRes, nil)
	c.JSON(http.StatusOK, successRes)
	return
}

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
