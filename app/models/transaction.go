package models

import (
	"gorm.io/datatypes"
	"martpay/app/enums"
	"martpay/app/request"
)

type Transactions struct {
	BaseModel
	UserId         uint           `json:"user_id" gorm:"not null"`
	TerminalId     uint           `json:"terminal_id" gorm:"not null"`
	TypeId         uint           `json:"type_id" gorm:"not null"`
	Amount         float64        `json:"amount" gorm:"not null"`
	Charge         float64        `json:"charge" gorm:"default:0.0"`
	TotalAmount    float64        `json:"total_amount" gorm:"not null"`
	Reference      string         `json:"reference" gorm:"unique;not null"`
	ResponseCode   *string        `json:"response_code" gorm:"type:varchar(4);"`
	Stan           *string        `json:"stan"`
	Recipient      *string        `json:"recipient"`
	BankName       *string        `json:"bank_name"`
	BankCode       *string        `json:"bank_code"`
	AccountName    *string        `json:"account_name"`
	Info           *string        `json:"info"`
	PowerToken     *string        `json:"power_token"`
	Status         enums.Status   `json:"status"`
	Channel        string         `json:"channel"`
	Provider       string         `json:"provider"`
	Meta           datatypes.JSON `json:"meta"`
	WalletDebited  bool           `json:"wallet_debited"`
	WalletCredited bool           `json:"wallet_credited"`
	Version        string         `json:"version"`
	Device         string         `json:"device"`
	Service        Service        `json:"service" gorm:"foreignKey:TypeId"`
}

func NewPendingTransaction(
	terminal *Terminal,
	service *Service,
	amount float64,
	charge float64,
	reference string,
	info string,
	provider string,
	recipient string,
	terminalInfo *request.TerminalInfo) *Transactions {

	return &Transactions{
		UserId:      terminal.UserId,
		TerminalId:  terminal.ID,
		TypeId:      service.ID,
		Amount:      amount,
		Charge:      charge,
		TotalAmount: charge + amount,
		Reference:   reference,
		Recipient:   &recipient,
		Status:      enums.PENDING,
		Info:        &info,
		Provider:    provider,
		Channel:     terminalInfo.CHANNEL,
		Version:     terminalInfo.VERSION,
		Device:      terminalInfo.DEVICE,
	}
}
