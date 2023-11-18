package usecase

import (
	"errors"

	"github.com/Anandhu4456/go-Ecommerce/pkg/domain"
	interfaces "github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
)

type paymentUsecase struct {
	paymentRepo interfaces.PaymentRepository
	userRepo    interfaces.UserRepository
}

// Constructor function
func NewPaymentUsecase(paymentRepo interfaces.PaymentRepository, userRepo interfaces.UserRepository) services.PaymentUsecase {
	return &paymentUsecase{
		paymentRepo: paymentRepo,
		userRepo:    userRepo,
	}
}

func (payU *paymentUsecase) AddNewPaymentMethod(paymentMethod string) error {
	if paymentMethod == "" {
		return errors.New("enter payment method")
	}
	if err := payU.paymentRepo.AddNewPaymentMethod(paymentMethod); err != nil {
		return err
	}
	return nil
}

func (payU *paymentUsecase) RemovePaymentMethod(paymentMethodID int) error {
	if paymentMethodID == 0 {
		return errors.New("enter method id")
	}
	if err := payU.paymentRepo.RemovePaymentMethod(paymentMethodID); err != nil {
		return err
	}
	return nil
}

func (payU *paymentUsecase) GetPaymentMethods() ([]domain.PaymentMethod, error) {
	paymentMethods, err := payU.paymentRepo.GetPaymentMethods()
	if err != nil {
		return []domain.PaymentMethod{}, err
	}
	return paymentMethods, nil
}
