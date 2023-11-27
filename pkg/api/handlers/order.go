package handlers

import (
	services"github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
)

type OrderHandler struct{
	orderUsecase services.OrderUsecase
}

// Constructor function

func NewOrderHandler(orderUsecase services.OrderUsecase)*OrderHandler{
	return &OrderHandler{
		orderUsecase: orderUsecase,
	}
}