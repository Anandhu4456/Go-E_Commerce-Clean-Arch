package usecase

import (
	"errors"

	"github.com/Anandhu4456/go-Ecommerce/pkg/config"
	"github.com/Anandhu4456/go-Ecommerce/pkg/helper"
	interfaces "github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
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

func (otU *otpUsecase) VerifyOTP(code models.VerifyData) (models.UserToken, error) {
	helper.TwilioSetup(otU.cfg.ACCOUNTSID, otU.cfg.AUTHTOKEN)
	if err := helper.TwiloVerifyOTP(otU.cfg.SERVICEID, code.Code, code.PhoneNumber); err != nil {
		return models.UserToken{}, errors.New("error while verifying OTP")
	}
	// getting user details to generate user token after verify OTP
	userDetails, err := otU.otpRepo.UserDetailsUsingPhone(code.PhoneNumber)
	if err != nil {
		return models.UserToken{}, err
	}
	tokenString, err := helper.GenerateUserToken(userDetails)
	if err != nil {
		return models.UserToken{}, err
	}
	return models.UserToken{
		Username: userDetails.Username,
		Token:    tokenString,
	}, nil
}
