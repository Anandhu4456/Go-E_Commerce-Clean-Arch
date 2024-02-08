package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userusecase services.UserUsecase
}

// Constructor function
func NewUserHandler(userUsecase services.UserUsecase) *UserHandler {
	return &UserHandler{
		userusecase: userUsecase,
	}
}

// @Summary		Add Address
// @Description	user can add their addresses
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			id	query	string	true	"id"
// @Param			address  body  models.AddAddress  true	"address"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/profile/address/add [post]
func (uH *UserHandler) AddAddress(c *gin.Context) {
	// id ,err:=helper.GetUserId(c)
	id, err := strconv.Atoi(c.Query("id"))
	fmt.Println("user id from add address handler ", id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check path parameter", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	var address models.AddAddress
	if err := c.BindJSON(&address); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	if err := uH.userusecase.AddAddress(id, address); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't add address", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully added address", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Change Password
// @Description	user can change their password
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			id	query	string	true	"id"
// @Param			changepassword  body  models.ChangePassword  true	"changepassword"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/profile/security/change-password [patch]
func (uH *UserHandler) ChangePassword(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "check path parameter", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	var changePass models.ChangePassword

	if err := c.BindJSON(&changePass); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := uH.userusecase.ChangePassword(userId, changePass.OldPassword, changePass.NewPassword, changePass.RePassword); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't change the password", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "password changed successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Edit User
// @Description	user can change their Details
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			id	query	string	true	"id"
// @Param			userData  body  models.EditUser true	"edit-user"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/profile/edit [patch]
func (uH *UserHandler) EditUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't get bad request", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	var userData models.EditUser
	if err := c.BindJSON(&userData); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := uH.userusecase.EditUser(userId, userData); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't change the user details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully changed user details", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Get Addresses
// @Description	user can get all their addresses
// @Tags			User
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Param			id	query	string	true	"id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/profile/address [get]
func (uH *UserHandler) GetAddresses(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "check id again", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	addresses, err := uH.userusecase.GetAddresses(userId)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully got all addresses", addresses, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Get Cart
// @Description	user can view their cart details
// @Tags			User
// @Produce		    json
// @Security		Bearer
// @Param			id	query	string	true	"id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/cart [get]
func (uH *UserHandler) GetCart(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "check path parameter", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	products, err := uH.userusecase.GetCart(userId)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't retrieve cart products", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully got all products in cart", products, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Get User Details
// @Description	user can get all their details
// @Tags			User
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Param			id	query	string	true	"id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/profile/details [get]
func (uH *UserHandler) GetUserDetails(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't get user id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	userDetails, err := uH.userusecase.GetUserDetails(userId)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't get user details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully got user details", userDetails, nil)
	c.JSON(http.StatusOK, successRes)
}

// Login is a handler for user login
// @Summary		User Login
// @Description	user can log in by giving their details
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			login  body  models.UserLogin  true	"login"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/login [post]
func (uH *UserHandler) Login(c *gin.Context) {
	var user models.UserLogin
	if err := c.BindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	userToken, err := uH.userusecase.Login(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "user couldn't login", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "user successfully logged in", userToken, nil)
	// c.SetCookie("Authorization",userToken.Token,3600,"/","yoursstore.online",true,false)
	c.SetCookie("Authorization", userToken.Token, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, successRes)
}

// @Summary		Remove from Cart
// @Description	user can remove products from their cart
// @Tags			User
// @Produce		    json
// @Param			id	query	string	true	"id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/cart/remove [delete]
func (uH *UserHandler) RemoveFromCart(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't find user id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	cartId, err := uH.userusecase.GetCartID(userId)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't get cart id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	invId, err := strconv.Atoi(c.Query("inventory_id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := uH.userusecase.RemoveFromCart(cartId, invId); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "remove from cart failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully removed from cart", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// Signup is a handler for user Registration
// @Summary		User Signup
// @Description	user can signup by giving their details
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			signup  body  models.UserDetails  true	"signup"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/signup [post]
func (uH *UserHandler) SignUp(c *gin.Context) {
	var user models.UserDetails

	if err := c.BindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	userToken, err := uH.userusecase.SignUp(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't signup user", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully signed up", userToken, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Add quantity in cart by one
// @Description	user can add 1 quantity of product to their cart
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param           id          query   string  true   "id"
// @Param			inventory 	query	string	true	"inventory id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/cart/updateQuantity/plus [post]
func (uH *UserHandler) UpdateQuantityAdd(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't get user id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	invId, err := strconv.Atoi(c.Query("inventory_id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := uH.userusecase.UpdateQuantityAdd(userId, invId); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't update quantity", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully added quantity", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Subtract quantity in cart by one
// @Description	user can subtract 1 quantity of product from their cart
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param           id          query   string  true    "id"
// @Param			inventory	query	string	true	"inventory id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/cart/updateQuantity/minus [post]
func (uH *UserHandler) UpdateQuantityLess(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "check path parameter", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	invId, err := strconv.Atoi(c.Query("inventory_id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := uH.userusecase.UpdateQuantityLess(userId, invId); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't subtract quantity", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully subtracted quantity", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// func (uH *UserHandler) GetWallet(c *gin.Context) {
// 	userId, err := helper.GetUserId(c)
// 	if err != nil {
// 		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't get user id", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errRes)
// 		return
// 	}
// 	page, err := strconv.Atoi(c.Query("page"))
// 	if err != nil {
// 		errRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errRes)
// 		return
// 	}
// 	limit, err := strconv.Atoi(c.Query("limit"))
// 	if err != nil {
// 		errRes := response.ClientResponse(http.StatusBadRequest, "limit number not in right format", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errRes)
// 		return
// 	}
// 	wallet, err := uH.userusecase.GetWallet(userId, page, limit)
// 	if err != nil {
// 		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't get wallet", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errRes)
// 		return
// 	}

// 	successRes := response.ClientResponse(http.StatusOK, "successfully get wallet", wallet, nil)
// 	c.JSON(http.StatusOK, successRes)
// }
