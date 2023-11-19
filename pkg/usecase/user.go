package usecase

import (
	interfaces"github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	services"github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
)
type userUsecase struct{
	userRepo interfaces.UserRepository
	offerRepo interfaces.OfferRepository
	walletRepo interfaces.WalletRepository
}

// Constructor function
func NewUserUsecase(userRepo interfaces.UserRepository,offerRepo interfaces.OfferRepository,walletRepo interfaces.WalletRepository)services.UserUsecase{
	return &userUsecase{
		userRepo: userRepo,
		offerRepo: offerRepo,
		walletRepo: walletRepo,
	}
}