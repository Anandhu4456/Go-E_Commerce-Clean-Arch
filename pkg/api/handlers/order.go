package handlers

import (
	"net/http"
	"strconv"

	"github.com/Anandhu4456/go-Ecommerce/pkg/helper"
	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderUsecase services.OrderUsecase
}

// Constructor function

func NewOrderHandler(orderUsecase services.OrderUsecase) *OrderHandler {
	return &OrderHandler{
		orderUsecase: orderUsecase,
	}
}

func (orH *OrderHandler) GetOrders(c *gin.Context) {
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
	id, err := helper.GetUserId(c)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "geting user id failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	orders, err := orH.orderUsecase.GetOrders(id, page, limit)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully got orders", orders, nil)
	c.JSON(http.StatusOK, successRes)
}

func (orH *OrderHandler) OrderItemsFromCart(c *gin.Context) {
	userId, err := helper.GetUserId(c)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't get user id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	coupStr := c.Query("coupon")
	coupon, err := strconv.Atoi(coupStr)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "conversion to integer failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	var order models.Order
	if err := c.BindJSON(&order); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	orderItemsString, err := orH.orderUsecase.OrderItemsFromCart(userId, order, coupon)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not make the order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully ordered", orderItemsString, nil)
	c.JSON(http.StatusOK, successRes)
}
