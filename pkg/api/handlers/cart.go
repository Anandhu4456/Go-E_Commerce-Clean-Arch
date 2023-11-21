package handlers

import (
	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
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
