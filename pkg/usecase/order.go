package usecase

import ("github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	services"github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"

)

type orderUsecase struct{
	orderRepo interfaces.OrderRepository
	userUsecase services.UserUsecase
	walletRepo interfaces.WalletRepository
	couponRepo interfaces.CouponRepository
}