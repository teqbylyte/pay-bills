package models

import (
	"gorm.io/datatypes"
	"time"
)

type Transactions struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	UserId         uint           `json:"user_id" gorm:"not null"`
	TerminalId     uint           `json:"terminal_id" gorm:"not null"`
	TypeId         uint           `json:"type_id" gorm:"not null"`
	Amount         float64        `json:"amount" gorm:"not null"`
	Charge         float64        `json:"charge" gorm:"default:0.0"`
	TotalAmount    float64        `json:"total_amount" gorm:"not null"`
	Reference      string         `json:"reference" gorm:"unique;not null"`
	ResponseCode   *string        `json:"response_code" gorm:"type:varchar(4);"`
	Stan           *string        `json:"stan"`
	BankName       *string        `json:"bank_name"`
	BankCode       *string        `json:"bank_code"`
	AccountNumber  *string        `json:"account_number"`
	AccountName    *string        `json:"account_name"`
	Info           *string        `json:"info"`
	PowerToken     *string        `json:"power_token"`
	Status         string         `json:"status"`
	Channel        string         `json:"channel"`
	Provider       string         `json:"provider"`
	Meta           datatypes.JSON `json:"meta"`
	WalletDebited  bool           `json:"wallet_debited"`
	WalletCredited bool           `json:"wallet_credited"`
	Version        string         `json:"version"`
	Device         string         `json:"device"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}
