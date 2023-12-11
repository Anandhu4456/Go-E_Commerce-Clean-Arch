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

// @Summary		Send OTP
// @Description	OTP login send otp
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			otp  body  models.OTPData true	"otp-data"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/otplogin [post]
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

// @Summary		Verify OTP
// @Description	OTP login verify otp
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			otp  body  models.VerifyData  true	"otp-verify"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/verifyotp [post]
func (otH *OtpHandler) VerifyOTP(c *gin.Context) {
	var code models.VerifyData
	if err := c.BindJSON(&code); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	users, err := otH.otpUsecase.VerifyOTP(code)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't verify OTP", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Verifyed OTP", users, nil)
	c.SetCookie("Authorization", users.Token, 3600, "", "", true, false)
	c.JSON(http.StatusOK, successRes)
}
