package handlers

import (
	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
)

type UserHandler struct {
	userusecase services.UserUsecase
}

// Constructor function
func NewUserHandler(userUsecase services.UserUsecase) *UserHandler {
	return &UserHandler{
		userusecase: userUsecase,
	}
}
