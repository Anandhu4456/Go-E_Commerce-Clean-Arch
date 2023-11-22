package handlers

import (
	services"github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
)

type WishlistHandler struct{
	wishlistUsecase services.WishlistUsecase
}

// Constructor function 
func NewWishlistHandler(wishlistUsecase services.WishlistUsecase)*WishlistHandler{
	return &WishlistHandler{
		wishlistUsecase: wishlistUsecase,
	}
}