package usecase

import (
	"errors"

	// "github.com/Anandhu4456/go-Ecommerce/pkg/domain"
	interfaces "github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
)

type CartUsecase struct {
	cartRepo       interfaces.CartRepository
	invRepo        interfaces.InventoryRespository
	userUsecase    services.UserUsecase
	paymentUsecase services.PaymentUsecase
}

// Constructor funciton

func NewCartUsecase(cartRepo interfaces.CartRepository, invRepo interfaces.InventoryRespository, userUsecase services.UserUsecase, paymentUsecase services.PaymentUsecase) *CartUsecase {
	return &CartUsecase{
		cartRepo:       cartRepo,
		invRepo:        invRepo,
		userUsecase:    userUsecase,
		paymentUsecase: paymentUsecase,
	}
}

func (cu *CartUsecase) AddToCart(user_id, inventory_id int) error {

	// check the product has quantity available
	stock, err := cu.invRepo.CheckStock(inventory_id)
	if err != nil {
		return errors.New("no stock")
	}
	if stock <= 0 {
		return errors.New("out of stock")
	}
	// Find user cart id
	cartId, err := cu.cartRepo.GetCartId(user_id)
	if err != nil {
		return errors.New("cart id not found")
	}
	// If user has no cart,create a cart
	if cartId == 0 {
		cartId, err = cu.cartRepo.CreateNewCart(user_id)
		if err != nil {
			return errors.New("cart creation failed")
		}
	}
	// Check if already added

	if cu.cartRepo.CheckIfInvAdded(inventory_id, cartId) {
		err := cu.cartRepo.AddQuantity(inventory_id, cartId)
		if err != nil {
			return err
		}
		return nil
	}
	// add product in line item
	err = cu.cartRepo.AddLineItems(inventory_id, cartId)
	if err != nil {
		return errors.New("product adding failed")
	}
	return nil
}

func (cu *CartUsecase) CheckOut(id int) (models.CheckOut, error) {

	// Getting address
	address, err := cu.cartRepo.GetAddresses(id)
	if err != nil {
		return models.CheckOut{}, errors.New("address not found")
	}
	products, err := cu.userUsecase.GetCart(id)
	if err != nil {
		return models.CheckOut{}, err
	}
	paymentMethod, err := cu.paymentUsecase.GetPaymentMethods()
	if err != nil {
		return models.CheckOut{}, err
	}
	var price, discount float64

	for _, v := range products.Values {
		discount += v.DiscountPrice
		price += v.Total
	}

	var checkOut models.CheckOut
	checkOut.CartId = products.Id
	checkOut.Addresses = address
	checkOut.Products = products.Values
	checkOut.PaymentMethods = paymentMethod
	checkOut.TotalPrice = price
	checkOut.DiscountPrice = discount

	return checkOut, nil
}
