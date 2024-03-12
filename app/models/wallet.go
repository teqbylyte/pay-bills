package models

import (
	"gorm.io/datatypes"
	"martpay/app/enums"
)

type Wallet struct {
	BaseModel
	Uwid          string         `json:"uwid" gorm:"not null"`
	UserId        uint           `json:"user_id" gorm:"not null"`
	AccountNumber string         `json:"account_number" gorm:"type:varchar(10);unique;not null"`
	Balance       float64        `json:"balance" gorm:"default:0.0"`
	Commission    float64        `json:"commission" gorm:"default:0.0"`
	Status        enums.State    `json:"status" gorm:"type:varchar(20);default:ACTIVE"`
	DisableDebit  bool           `json:"disable_debit" gorm:"default:false"`
	Meta          datatypes.JSON `json:"meta"`
}
