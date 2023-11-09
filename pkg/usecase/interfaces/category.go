package interfaces

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/domain"
)

type CategoryUsecase interface {
	AddCategory(category string) (domain.Category, error)
	UpdateCategory(currrent, new string) (domain.Category, error)
	DeleteCategory(categoryId string) error
	GetCategories(page, limit int) ([]domain.Category, error)
}
