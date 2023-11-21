package handlers

import (
	services"github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
)

type AdminHandler struct{
	adminUsecase services.AdminUsecase 
}
// constructor function
func NewAdminHandler(adminUsecase services.AdminUsecase)*AdminHandler{
	return &AdminHandler{
		adminUsecase: adminUsecase,
	}
}