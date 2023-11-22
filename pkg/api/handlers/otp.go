package handlers

import (
	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
)

type OtpHandler struct {
	otpUsecase services.OtpUsecase
}

// Constrctor function
func NewOtpHandler(otpUsecase services.OtpUsecase) *OtpHandler {
	return &OtpHandler{
		otpUsecase: otpUsecase,
	}
}
