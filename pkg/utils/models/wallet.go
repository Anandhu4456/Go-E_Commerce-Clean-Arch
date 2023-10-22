package models

import "github.com/Anandhu4456/go-Ecommerce/pkg/domain"

type Wallet struct {
	Balance int                    `json:"balance"`
	History []domain.WalletHistory `json:"history"`
}
