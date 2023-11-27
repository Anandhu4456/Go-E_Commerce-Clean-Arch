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
	coupon := c.Query("coupon")

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

func (orH *OrderHandler) CancelOrder(c *gin.Context) {
	userId, err := helper.GetUserId(c)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't get user id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	orderIdStr := c.Query("order_id")
	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	if err := orH.orderUsecase.CancelOrder(userId, orderId); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "cancel order failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully canceled the order", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (orH *OrderHandler) EditOrderStatus(c *gin.Context) {
	status := c.Query("status")
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := orH.orderUsecase.EditOrderStatus(status, id); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "edit order status failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully edited order status", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (orH *OrderHandler) MarkAsPaid(c *gin.Context) {
	orderId, err := strconv.Atoi(c.Query("order_id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := orH.orderUsecase.MarkAsPaid(orderId); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "mark as paid failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully edited payment status", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (orH *OrderHandler) AdminOrders(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "limit number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	status := c.Query("status")
	orders, err := orH.orderUsecase.AdminOrders(page, limit, status)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully got all records", orders, nil)
	c.JSON(http.StatusOK, successRes)
}

func (orH *OrderHandler) AdminSalesDailyReport(c *gin.Context) {
	salesReport, err := orH.orderUsecase.DailyOrders()
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully got all records", salesReport, nil)
	c.JSON(http.StatusOK, successRes)
}

func (orH *OrderHandler) AdminSalesWeeklyReports(c *gin.Context) {
	salesReport, err := orH.orderUsecase.WeeklyOrders()
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully got all records", salesReport, nil)
	c.JSON(http.StatusOK, successRes)
}
