package handlers

import (
	"net/http"

	"github.com/Anandhu4456/go-Ecommerce/pkg/domain"
	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type CouponHandler struct {
	couponUsecase services.CouponUsecase
}

// Constructor function
func NewCouponHandler(couponUsecase services.CouponUsecase) *CouponHandler {
	return &CouponHandler{
		couponUsecase: couponUsecase,
	}
}

func (coupH *CouponHandler) CreateNewCoupon(c *gin.Context) {
	var coupon domain.Coupon

	if err := c.BindJSON(&coupon); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	if err := coupH.couponUsecase.Addcoupon(coupon); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add the Coupon", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully added the coupon", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
