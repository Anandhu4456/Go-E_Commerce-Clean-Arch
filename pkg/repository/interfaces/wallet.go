package interfaces

import "github.com/Anandhu4456/go-Ecommerce/pkg/domain"

type WalletRepository interface{
	CreateNewWallet(userId int)(int,error)
	CreditToUserWallet(amount float64, walletId int)error
	FindUserIdFromOrderId(id int)(int,error)
	FindWalletIdFromUserId(userId int)(int,error)
	GetBalance(walletId int)(int,error)
	GetHistory(walletId,page,limit int)([]domain.WalletHistory,error)
	AddHistory(amount,walletId int,purpose string)error
	PayFromWallet(userId,orderId int,price float64)(float64,error)
}