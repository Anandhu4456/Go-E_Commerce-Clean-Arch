package usecase

import (
	"errors"

	interfaces "github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
)

type offerUsecase struct {
	offerRepo interfaces.OfferRepository
}

// constructor function

func NewOfferUsecase(offerRepo interfaces.OfferRepository) services.OfferUsecase {
	return &offerUsecase{
		offerRepo: offerRepo,
	}
}

func (offU *offerUsecase) AddNewOffer(model models.CreateOffer) error {
	if err := offU.offerRepo.AddNewOffer(model); err != nil {
		return errors.New("adding offer failed")
	}
	return nil
}

func (offU *offerUsecase) MakeOfferExpire(catID int) error {
	if err := offU.offerRepo.MakeOfferExpired(catID); err != nil {
		return err
	}
	return nil
}
