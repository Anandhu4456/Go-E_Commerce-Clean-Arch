package usecase

import (
	interfaces"github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	services"github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
)

type offerUsecase struct{
	offerRepo interfaces.OfferRepository
}
// constructor function

func NewOfferUsecase(offerRepo interfaces.OfferRepository)services.OfferUsecase{
	return &offerUsecase{
		offerRepo: offerRepo,
	}
}