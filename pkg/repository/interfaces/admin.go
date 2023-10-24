package interfaces

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/domain"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
)

type AdminRepository interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error)
	GetUserById(id string) (domain.User, error)
	UpdateBlockUserById(user domain.User) error
	GetUsers(page, limit int) ([]models.UserDetailsAtAdmin, error)
}
