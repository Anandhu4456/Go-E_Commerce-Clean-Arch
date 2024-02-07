package handlers

import (
	"net/http"
	"strconv"

	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
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

// @Summary		Add Coupon
// @Description	Admin can add new coupons
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			coupon	body	models.Coupon	true	"coupon"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/coupons/create [post]
func (coupH *CouponHandler) CreateNewCoupon(c *gin.Context) {
	var coupon models.Coupon

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

// @Summary		Make Coupon invalid
// @Description	Admin can make the coupons as invalid so that users cannot use that particular coupon
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			id	query	string	true	"id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/coupons/expire [post]
func (coupH *CouponHandler) MakeCouponInvalid(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "field provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := coupH.couponUsecase.MakeCouponInvalid(id); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "make coupon invalid failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully made coupon as invalid", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		List Coupons
// @Description	Admin can view the list of  Coupons
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			page	query  string 	true	"page"
// @Param			limit	query  string 	true	"limit"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/coupons [get]
func (coupH *CouponHandler) Coupons(c *gin.Context) {

	coupons, err := coupH.couponUsecase.GetCoupons()
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't get coupons", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "coupons get successfully", coupons, nil)
	c.JSON(http.StatusOK, successRes)

}
