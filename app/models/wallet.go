package models

import "gorm.io/datatypes"

type Wallet struct {
	BaseModel
	Uwid string `json:"uwid" gorm:"not null"`
	BelongsToUser
	AccountNumber string         `json:"account_number"`
	Balance       float64        `json:"balance"`
	Commission    float64        `json:"commission"`
	Status        string         `json:"status"`
	DisableDebit  bool           `json:"disable_debit"`
	Meta          datatypes.JSON `json:"meta"`
}
