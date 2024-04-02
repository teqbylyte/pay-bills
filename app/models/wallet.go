package model

import (
	"errors"
	"fmt"
	"gorm.io/datatypes"
	"pay-bills/app/enums"
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

func (w Wallet) IsActive() bool {
	return w.Status == enums.ACTIVE
}

// AllowsTransaction - Ensure wallet can perform transaction.
func (w Wallet) AllowsTransaction(amount float64, forDebit bool, allowNegative bool) error {
	fmt.Println(w.ID)
	if !w.IsActive() {
		return errors.New(fmt.Sprintf("Your wallet is %s", w.Status))
	}

	if forDebit && !allowNegative && amount > w.Balance {
		return errors.New("Insufficient balance")
	}

	return nil
}
