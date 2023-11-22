package handlers

import (
	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
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
