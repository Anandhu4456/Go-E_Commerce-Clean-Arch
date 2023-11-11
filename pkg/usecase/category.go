package usecase

import (
	"errors"

	"github.com/Anandhu4456/go-Ecommerce/pkg/domain"
	"github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
)

type categoryUsecase struct {
	repo interfaces.CategoryRepository
}

// constructor function

func NewCategoryUsecase(repo interfaces.CategoryRepository) services.CategoryUsecase {
	return categoryUsecase{
		repo: repo,
	}
}

func (catU *categoryUsecase) AddCategory(category string) (domain.Category, error) {
	productResponse, err := catU.repo.AddCategory(category)
	if err != nil {
		return domain.Category{}, err
	}
	return productResponse, nil
}

func (catU *categoryUsecase) UpdateCategory(currrent, new string) (domain.Category, error) {
	result, err := catU.repo.CheckCategory(currrent)
	if err != nil {
		return domain.Category{}, err
	}
	if !result {
		return domain.Category{}, errors.New("no category as you mentioned")
	}
	newCat, err := catU.UpdateCategory(currrent, new)
	if err != nil {
		return domain.Category{}, err
	}
	return newCat, nil
}
