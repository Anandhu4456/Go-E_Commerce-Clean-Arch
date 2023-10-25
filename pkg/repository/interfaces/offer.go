package interfaces

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/domain"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
)

type OfferRepository interface {
	AddNewOffer(models.CreateOffer) error
	MakeOfferExpired(cartId int) error
	FindDiscountPercentage(cartId int) (int, error)
	GetOffers(page, limit int) ([]domain.Offer, error)
}
