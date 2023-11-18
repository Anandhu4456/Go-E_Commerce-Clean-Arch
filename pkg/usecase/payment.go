package usecase

import (
	interfaces"github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	services"github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
)

type paymentUsecase struct{
	paymentRepo interfaces.PaymentRepository
	userRepo interfaces.UserRepository
}
// Constructor function
func NewPaymentUsecase(paymentRepo interfaces.PaymentRepository,userRepo interfaces.UserRepository)services.PaymentUsecase{
	return &paymentUsecase{
		paymentRepo: paymentRepo,
		userRepo: userRepo,
	}
}