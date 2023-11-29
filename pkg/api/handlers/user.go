package handlers

import (
	"net/http"

	"github.com/Anandhu4456/go-Ecommerce/pkg/helper"
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

func (uH *UserHandler) AddAddress(c *gin.Context) {
	userId, err := helper.GetUserId(c)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't get user id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	var address models.AddAddress
	if err := c.BindJSON(&address); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := uH.userusecase.AddAddress(userId, address); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't add address", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully added address", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (uH *UserHandler) ChangePassword(c *gin.Context) {
	userId, err := helper.GetUserId(c)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't find user id", nil, err.Error())
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

func (uH *UserHandler) EditUser(c *gin.Context) {
	userId, err := helper.GetUserId(c)
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
}
