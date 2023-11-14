package usecase

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/domain"
	"github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
)

type orderUsecase struct {
	orderRepo   interfaces.OrderRepository
	userUsecase services.UserUsecase
	walletRepo  interfaces.WalletRepository
	couponRepo  interfaces.CouponRepository
}

func NewOrderUsecase(orderRepo interfaces.OrderRepository, userUsecase services.UserUsecase, walletRepo interfaces.WalletRepository, couponRepo interfaces.CouponRepository) *orderUsecase {
	return &orderUsecase{
		orderRepo:   orderRepo,
		userUsecase: userUsecase,
		walletRepo:  walletRepo,
		couponRepo:  couponRepo,
	}
}

func (orU *orderUsecase) GetOrders(id, page, limit int) ([]domain.Order, error) {
	orders, err := orU.orderRepo.GetOrders(id, page, limit)
	if err != nil {
		return []domain.Order{}, err
	}
	return orders, nil
}
