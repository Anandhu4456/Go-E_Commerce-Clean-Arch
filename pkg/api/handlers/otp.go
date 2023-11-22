package handlers

import (
	"net/http"

	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/response"
	"github.com/gin-gonic/gin"
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

func (otH *OtpHandler) SendOTP(c *gin.Context) {
	var phone models.OTPData

	if err := c.BindJSON(&phone); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := otH.otpUsecase.SendOTP(phone.PhoneNumber); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Could not sent OTP", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "OTP send successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
