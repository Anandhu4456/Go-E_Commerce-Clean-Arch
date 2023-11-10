package usecase

import (
	interfaces "github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
)

type cartUsecase struct {
	cartRepo       interfaces.CartRepository
	invRepo        interfaces.InventoryRespository
	userUsecase    services.UserUsecase
	paymentUsecase services.PaymentUsecase
}
