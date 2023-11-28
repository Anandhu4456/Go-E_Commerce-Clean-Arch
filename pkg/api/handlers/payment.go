package handlers

import (
	"net/http"

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
