package usecase

import ("github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
)

type categoryUsecase struct{
	repo interfaces.CategoryRepository
}

// constructor function

func NewCategoryUsecase(repo interfaces.CategoryRepository)services.CategoryUsecase{
	return categoryUsecase{
		repo:repo,
	}
}