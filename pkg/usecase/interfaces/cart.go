package interfaces

import "github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"

type CartUsecase interface {
	AddToCart(user_id, inventory_id int) error
	CheckOut(id int) (models.CheckOut, error)
}
