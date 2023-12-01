package handlers

import (
	"net/http"
	"strconv"

	"github.com/Anandhu4456/go-Ecommerce/pkg/helper"
	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	cartUsecase services.CartUsecase
}

// Constructor function
func NewCartHandler(usecase services.CartUsecase) *CartHandler {
	return &CartHandler{
		cartUsecase: usecase,
	}
}

func (ch *CartHandler) AddtoCart(c *gin.Context) {
	userId, err := helper.GetUserId(c)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not get userID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	inventoryID, err := strconv.Atoi(c.Query("inventory"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check parameters properly", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := ch.cartUsecase.AddToCart(userId, inventoryID); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add the Cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added To cart", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (ch *CartHandler) CheckOut(c *gin.Context) {
	userId, err := helper.GetUserId(c)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not get userID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	products, err := ch.cartUsecase.CheckOut(userId)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not open checkout", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", products, nil)
	c.JSON(http.StatusOK, successRes)
}
