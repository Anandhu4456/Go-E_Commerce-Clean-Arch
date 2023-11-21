package usecase

import (
	interfaces "github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
)

type wishlistUsecase struct {
	wishRepo interfaces.WishlistRepository
}

// Constructor function
func NewWishlistUsecase(wishRepo interfaces.WishlistRepository) services.WishlistUsecase {
	return &wishlistUsecase{
		wishRepo: wishRepo,
	}
}

func (wlU *wishlistUsecase) GetWishlistID(userID int) (int, error) {
	wishlistId, err := wlU.wishRepo.GetWishlistId(userID)
	if err != nil {
		return 0, err
	}
	return wishlistId, nil
}
