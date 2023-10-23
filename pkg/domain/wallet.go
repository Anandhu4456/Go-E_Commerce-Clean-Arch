package domain

import "time"

type Wallet struct {
	ID     int     `json:"id" gorm:"unique;not null"`
	UserID int     `json:"user_id"`
	User   User    `json:"-" gorm:"foreignkey:UserID"`
	Amount float64 `json:"amount" gorm:"default:0"`
}

type WalletHistory struct {
	WalletID int       `json:"walletID"`
	Wallet   Wallet    `json:"-" gorm:"foreignkey:WalletID"`
	Amount   int       `json:"amount"`
	Purpose  string    `json:"purpose"`
	Time     time.Time `json:"time"`
}
