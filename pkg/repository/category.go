package repository

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type categoryRepository struct{
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB)interfaces.CategoryRepository{
	return &categoryRepository{
		DB:db,
	}
}