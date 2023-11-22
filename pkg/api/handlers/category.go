package handlers

import (
	services"github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
)

type CategoryHandler struct{
	CategoryUsecase services.CategoryUsecase
}
// Constructor function

func NewCategoryHandler(categoryUsecase services.CategoryUsecase)*CategoryHandler{
	return &CategoryHandler{
		CategoryUsecase: categoryUsecase,
	}
}