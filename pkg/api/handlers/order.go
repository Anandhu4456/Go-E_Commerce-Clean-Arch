package handlers

import (
	"net/http"
	"os"
	"strconv"

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

// @Summary		Get Orders
// @Description	user can view the details of the orders
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param            id     query  string   true   "id"
// @Param			page	query  string 	true	"page"
// @Param			limit	query  string 	true	"limit"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/profile/orders [get]
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
	id, err := strconv.Atoi(c.Query("id"))
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

// @Summary		Order Now
// @Description	user can order the items that currently in cart
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			coupon	query	string	true	"coupon"
// @Param			order	body	models.Order	true	"order"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/check-out/order [post]
func (orH *OrderHandler) OrderItemsFromCart(c *gin.Context) {

	var order models.Order
	if err := c.BindJSON(&order); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := orH.orderUsecase.OrderItemsFromCart(order.UserID, order.AddressID, order.PaymentMethodID, order.CouponID); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not make the order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully ordered", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Order Cancel
// @Description	user can cancel the orders
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			orderid  query  string  true	"order id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/profile/orders/cancel [post]
func (orH *OrderHandler) CancelOrder(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("id"))
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

// @Summary		Update Order Status
// @Description	Admin can change the status of the order
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			id  query  string  true	"id"
// @Param			status  query  string  true	"status"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/orders/edit/status [patch]
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

// @Summary		Update Payment Status
// @Description	Admin can change the status of the payment
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			orderID  query  string  true	"order id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/orders/edit/mark-as-paid [patch]
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

// @Summary		Admin Orders
// @Description	Admin can view the orders according to status
// @Tags			Admin
// @Produce		    json
// @Param			page	query  string 	true	"page"
// @Param			limit	query  string 	true	"limit"
// @Param			status	query  string	true	"status"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/orders [get]
func (orH *OrderHandler) AdminOrders(c *gin.Context) {

	orders, err := orH.orderUsecase.AdminOrders()
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully got all records", orders, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Admin Sales Report
// @Description	Admin can view the daily sales Report
// @Tags			Admin
// @Produce		    json
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/sales/daily [get]
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

// @Summary		Admin Sales Report
// @Description	Admin can view the weekly sales Report
// @Tags			Admin
// @Produce		    json
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/sales/weekly [get]
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

// @Summary		Admin Sales Report
// @Description	Admin can view the weekly sales Report
// @Tags			Admin
// @Produce		    json
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/sales/monthly [get]
func (orH *OrderHandler) AdminSalesMonthlyReport(c *gin.Context) {
	salesReport, err := orH.orderUsecase.MonthlyOrders()
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully got all records", salesReport, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Admin Sales Report
// @Description	Admin can view the weekly sales Report
// @Tags			Admin
// @Produce		    json
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/sales/annual [get]
func (orH *OrderHandler) AdminSalesAnnualReport(c *gin.Context) {
	salesReport, err := orH.orderUsecase.AnnualOrders()
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully got all records", salesReport, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Admin Sales Report
// @Description	Admin can view the weekly sales Report
// @Tags			Admin
// @Produce		    json
// @Param			customDates  body  models.CustomDates  true	"custom dates"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/sales/custom [post]
func (orH *OrderHandler) AdminSaleCustomReport(c *gin.Context) {
	var dates models.CustomDates
	if err := c.BindJSON(&dates); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	salesReport, err := orH.orderUsecase.CustomDateOrders(dates)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully got all records", salesReport, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Return Order
// @Description	user can return the ordered products which is already delivered and then get the amount fot that particular purchase back in their wallet
// @Tags			User
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Param			id  query  string  true	"id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/profile/orders/return [post]
func (orH *OrderHandler) ReturnOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := orH.orderUsecase.ReturnOrder(id); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fileds provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "return order successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Download Invoice PDF
// @Description Download the invoice PDF file
// @Tags			User
// @Security		Bearer
// @Produce octet-stream
// @Success 200 {file} application/pdf
// @Router /users/check-out/order/download-invoice  [get]
func (orH *OrderHandler) DownloadInvoice(c *gin.Context) {
	// Set the appropriate header for the file download
	c.Header("Content-Disposition", "attachment; filename=yoursstore_invoice.pdf")
	c.Header("Content-Type", "application/pdf")

	// Read the pdf file and write it to the response
	pdfData, err := os.ReadFile("yoursstore_invoice.pdf")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read pdf file"})
		return
	}
	c.Data(http.StatusOK, "application/pdf", pdfData)
}
