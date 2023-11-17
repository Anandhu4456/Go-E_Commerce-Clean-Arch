package usecase

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/config"
	interfaces "github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
)

type otpUsecase struct {
	cfg     config.Config
	otpRepo interfaces.OtpRepository
}

// constructor function
func NewOtpUsecase(cfg config.Config, otpRepo interfaces.OtpRepository) services.OtpUsecase {
	return &otpUsecase{
		cfg:     cfg,
		otpRepo: otpRepo,
	}
}
