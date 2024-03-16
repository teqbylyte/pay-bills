package dto

import (
	"gorm.io/datatypes"
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

type LoanDto struct {
	Id            uint           `json:"id"`
	Amount        float64        `json:"amount"`
	Charge        float64        `json:"charge"`
	AmountOwed    float64        `json:"amount_owed"`
	Items         datatypes.JSON `json:"items"`
	Status        string         `json:"status"`
	Info          *string        `json:"info"`
	DeclineReason *string        `json:"decline_reason"`
	CreatedAt     time.Time      `json:"created_at"`
}
