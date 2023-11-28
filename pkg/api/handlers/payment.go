package handlers

import (
	"net/http"
	"strconv"

	"github.com/Anandhu4456/go-Ecommerce/pkg/helper"
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

func (payH *PaymentHandler) MakePamentRazorPay(c *gin.Context) {
	orderId := c.Query("id")
	userId, err := helper.GetUserId(c)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't get user id")
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
