package handlers

import (
	"net/http"

	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminUsecase services.AdminUsecase
}

// constructor function
func NewAdminHandler(adminUsecase services.AdminUsecase) *AdminHandler {
	return &AdminHandler{
		adminUsecase: adminUsecase,
	}
}

func (ah *AdminHandler) LoginHandler(c *gin.Context) {
	// login handler for the admin
	var adminDetails models.AdminLogin

	if err := c.BindJSON(&adminDetails); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	admin, err := ah.adminUsecase.LoginHandler(adminDetails)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "can't authenticate admin", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", admin.Token, 3600, "/", "", true, false)

	successRes := response.ClientResponse(http.StatusOK, "Admin authenticated successfully", admin, nil)
	c.JSON(http.StatusOK, successRes)
}

func (ah *AdminHandler) BlockUser(c *gin.Context) {
	id := c.Query("id")
	err := ah.adminUsecase.BlockUser(id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "cant block", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "blocked the user", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
