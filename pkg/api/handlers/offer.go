package handlers

import (
	"net/http"

	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type OfferHandler struct {
	offerUsecase services.OfferUsecase
}

// Constructor function

func NewOfferHandler(offerUsecase services.OfferUsecase) *OfferHandler {
	return &OfferHandler{
		offerUsecase: offerUsecase,
	}
}

func (offH *OfferHandler) AddOffer(c *gin.Context) {
	var offer models.CreateOffer

	if err := c.BindJSON(&offer); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := offH.offerUsecase.AddNewOffer(offer); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "offer adding failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "offer added successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
