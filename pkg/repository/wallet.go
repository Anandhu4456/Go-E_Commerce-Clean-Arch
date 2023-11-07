package repository

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type walletRepository struct{
	DB *gorm.DB
}
// constructor function
func NewWalletRepository(DB *gorm.DB)interfaces.WalletRepository{
	return &walletRepository{
		DB:DB,
	}
}