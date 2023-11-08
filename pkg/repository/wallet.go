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

func (wr *walletRepository) FindWalletIdFromUserId(userId int) (int, error) {
	var walletCount int
	if err := wr.DB.Raw("SELECT COUNT(*)FROM wallets WHERE user_id=?", userId).Scan(&walletCount).Error; err != nil {
		return 0, errors.New("wallet not found")
	}
	var walletId int
	if walletCount > 0 {
		err := wr.DB.Raw("SELECT id FROM wallets WHERE user_id=?", userId).Scan(&walletId).Error
		if err != nil {
			return 0, errors.New("wallet id not found")
		}
	}
	return walletId, nil
}
