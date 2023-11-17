package usecase

import (
	"errors"

	"github.com/Anandhu4456/go-Ecommerce/pkg/config"
	"github.com/Anandhu4456/go-Ecommerce/pkg/helper"
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

func (otU *otpUsecase) SendOTP(phone string) error {
	ok := otU.otpRepo.FindUserByMobileNumber(phone)
	if !ok {
		return errors.New("phone number not found")
	}
	helper.TwilioSetup(otU.cfg.ACCOUNTSID, otU.cfg.AUTHTOKEN)
	_, err := helper.TwilioSendOTP(phone, otU.cfg.SERVICEID)
	if err != nil {
		return errors.New("error occured while generating OTP")
	}
	return nil
}
