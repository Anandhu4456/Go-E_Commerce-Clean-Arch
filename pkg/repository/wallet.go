package repository

import (
	"errors"

	"github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type walletRepository struct {
	DB *gorm.DB
}

// constructor function
func NewWalletRepository(DB *gorm.DB) interfaces.WalletRepository {
	return &walletRepository{
		DB: DB,
	}
}

func (wr *walletRepository) CreditToUserWallet(amount float64, walletId int) error {
	if err := wr.DB.Exec("UPDATE wallets SET amount=amount+$1 WHERE wallet_id=$2", amount, walletId).Error; err != nil {
		return errors.New("amount adding to wallet failed")
	}
	return nil
}

func (wr *walletRepository) FindUserIdFromOrderId(id int) (int, error) {
	var userId int
	err := wr.DB.Raw("SELECT user_id FROM orders WHERE order_id=?", id).Scan(&userId).Error
	if err != nil {
		return 0, errors.New("user id not found")
	}
	return userId, nil
}
