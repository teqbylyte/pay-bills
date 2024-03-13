package dto

import (
	"time"
)

type TransactionDto struct {
	Id             int       `json:"id"`
	Amount         int       `json:"amount"`
	Charge         float64   `json:"charge"`
	TotalAmount    float64   `json:"total_amount"`
	Reference      string    `json:"reference"`
	Status         string    `json:"status"`
	Info           string    `json:"info"`
	BankName       *string   `json:"bank_name"`
	AccountNumber  *string   `json:"account_number"`
	PowerToken     *string   `json:"power_token"`
	WalletCredited *bool     `json:"wallet_credited"`
	WalletDebited  *bool     `json:"wallet_debited"`
	CreatedAt      time.Time `json:"created_at"`
}
