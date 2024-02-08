package handlers

import (
	"net/http"
	"strconv"

	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	paymentUsecase services.PaymentUsecase
}

// Constructor function

func NewPaymentHandler(payUsecase services.PaymentUsecase) *PaymentHandler {
	return &PaymentHandler{
		paymentUsecase: payUsecase,
	}
}

// @Summary		Add new payment method
// @Description	admin can add a new payment method
// @Tags			Admin
// @Produce		    json
// @Param			paymentMethod	query  string 	true	"Payment Method"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/paymentmethods/add [post]
func (payH *PaymentHandler) AddNewPaymentMethod(c *gin.Context) {
	method := c.Query("payment_method")

	if err := payH.paymentUsecase.AddNewPaymentMethod(method); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't add new payment method", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully added payment method", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Remove payment method
// @Description	admin can remove a  payment method
// @Tags			Admin
// @Produce		    json
// @Param			paymentMethodID	query  int 	true	"Payment Method ID"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/paymentmethods/remove [delete]
func (payH *PaymentHandler) RemovePaymentMethod(c *gin.Context) {
	methodId, err := strconv.Atoi(c.Query("payment_method_id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := payH.paymentUsecase.RemovePaymentMethod(methodId); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "payment method removal failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "payment method removed successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Get payment methods
// @Description	admin can get all  payment methods
// @Tags			Admin
// @Produce		    json
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/paymentmethods [get]
func (payH *PaymentHandler) GetPaymentMethods(c *gin.Context) {
	paymentMethods, err := payH.paymentUsecase.GetPaymentMethods()

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't get payment methods", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully collected payment methods", paymentMethods, nil)
	c.JSON(http.StatusOK, successRes)
}

func (payH *PaymentHandler) MakePaymentRazorPay(c *gin.Context) {
	orderId := c.Query("id")
	userId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "check path paremeter(user id)", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	payDetails, err := payH.paymentUsecase.MakePaymentRazorPay(orderId, userId)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't generate order details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	c.HTML(http.StatusOK, "razorpay.html", payDetails)
}

func (payH *PaymentHandler) VerifyPayment(c *gin.Context) {
	paymentId := c.Query("payment_id")
	razorId := c.Query("razor_id")
	orderId := c.Query("order_id")

	if err := payH.paymentUsecase.VerifyPayment(paymentId, razorId, orderId); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't update payment details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully updated payment details", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
