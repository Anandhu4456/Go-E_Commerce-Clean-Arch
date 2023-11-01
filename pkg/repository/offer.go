package repository

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/domain"
	"github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
	"gorm.io/gorm"
)

type offerRepository struct {
	DB *gorm.DB
}

// constructor function

func NewOfferRepository(DB *gorm.DB) interfaces.OfferRepository {
	return &offerRepository{
		DB: DB,
	}
}

func (or *offerRepository) AddNewOffer(offer models.CreateOffer) error {
	err := or.DB.Exec("INSERT INTO offers(category_id,discount) VALUES (?,?)", offer.CategoryID, offer.Discount).Error
	if err != nil {
		return err
	}
	return nil
}

func (or *offerRepository) MakeOfferExpired(categoryId int) error {
	err := or.DB.Exec("UPDATE offers SET valid=false WHERE id=$1", categoryId).Error
	if err != nil {
		return err
	}
	return nil
}

func (or *offerRepository) FindDiscountPercentage(categoryId int) (int, error) {
	var percentage int

	err := or.DB.Raw("SELECT discount_rate FROM offers WHERE category_id=$1 and valid=true", categoryId).Scan(&percentage).Error
	if err != nil {
		return 0, err
	}
	return percentage, nil
}

func (or *offerRepository) GetOffers(page, limit int) ([]domain.Offer, error) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	var getOffers []domain.Offer

	if err := or.DB.Raw("SELECT id,category_id,discount_rate,valid FROM offers limit ? offset ?", limit, offset).Scan(&getOffers).Error; err != nil {
		return []domain.Offer{}, err
	}
	return getOffers, nil
}
