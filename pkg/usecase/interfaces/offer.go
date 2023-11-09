package interfaces

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/domain"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
)

type OfferUsecase interface {
	AddNewOffer(model models.CreateOffer) error
	MakeOfferExpire(catID int) error
	GetOffers(page, limit int) ([]domain.Offer, error)
}
