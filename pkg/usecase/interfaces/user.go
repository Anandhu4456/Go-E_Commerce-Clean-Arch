package interfaces

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/domain"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
)

type UserUsecase interface {
	Login(user models.UserLogin) (models.UserToken, error)
	SignUp(user models.UserDetails) (models.UserToken, error)
	AddAddress(id int, address models.AddAddress) error
	GetAddresses(id int) ([]domain.Address, error)
	GetUserDetails(id int) (models.UserResponse, error)

	ChangePassword(id int, old string, password string, repassword string) error
	EditUser(id int, userData models.EditUser) error

	GetCartID(userID int) (int, error)
	GetCart(id, page, limit int) ([]models.GetCart, error)
	RemoveFromCart(id int, inventoryID int) error
	ClearCart(cartID int) error
	UpdateQuantityAdd(id, inv_id int) error
	UpdateQuantityLess(id, inv_id int) error

	GetWallet(id, page, limit int) (models.Wallet, error)
}
