package handlers

import (
	services"github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
)

type PaymentHandler struct{
	paymentUsecase services.PaymentUsecase
}

// Constructor function

func NewPaymentHandler(payUsecase services.PaymentUsecase)*PaymentHandler{
	return &PaymentHandler{
		paymentUsecase: payUsecase,
	}
}