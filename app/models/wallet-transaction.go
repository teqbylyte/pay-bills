package models

import "martpay/app/enums"

type WalletTransaction struct {
	BaseModel
	ProductId   uint            `json:"product_id" gorm:"not null"`
	WalletId    uint            `json:"wallet_id" gorm:"not null"`
	Reference   string          `json:"reference" gorm:"not null"`
	Amount      float64         `json:"amount" gorm:"not null"`
	PrevBalance float64         `json:"prev_balance" gorm:"not null"`
	NewBalance  float64         `json:"new_balance" gorm:"not null"`
	Status      enums.Status    `json:"status" gorm:"not null"`
	Action      enums.ACTION    `json:"action" gorm:"not null"`
	Type        enums.TransType `json:"type" gorm:"not null"`
	Info        *string         `json:"info"`
}
