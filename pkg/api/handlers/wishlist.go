package handlers

import (
	"net/http"
	"strconv"

	"github.com/Anandhu4456/go-Ecommerce/pkg/helper"
	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type WishlistHandler struct {
	wishlistUsecase services.WishlistUsecase
}

// Constructor function
func NewWishlistHandler(wishlistUsecase services.WishlistUsecase) *WishlistHandler {
	return &WishlistHandler{
		wishlistUsecase: wishlistUsecase,
	}
}

func (wiH *WishlistHandler) AddtoWishlist(c *gin.Context) {
	userId, err := helper.GetUserId(c)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't get user id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	inventoryId, err := strconv.Atoi(c.Query("inventory"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "check parameters properly", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := wiH.wishlistUsecase.AddToWishlist(userId, inventoryId); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Could not add the wishlsit", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "added to wishlit", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (wiH *WishlistHandler) RemoveFromWishlist(c *gin.Context) {
	id, err := helper.GetUserId(c)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not get user id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	wishlistId, err := wiH.wishlistUsecase.GetWishlistID(id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "check parameters correctly", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	inv, err := strconv.Atoi(c.Query("inventory"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "check parameters correctly", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := wiH.wishlistUsecase.RemoveFromWishlist(wishlistId, inv); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "remove from wishlist failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "removed from wishlist", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
