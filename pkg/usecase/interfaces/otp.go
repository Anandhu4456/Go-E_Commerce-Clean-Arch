package interfaces

import "github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"

type OtpUsecase interface {
	VerifyOTP(code models.VerifyData) (models.UserToken, error)
	SendOTP(phone string) error
}
