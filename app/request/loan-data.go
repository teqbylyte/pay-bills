package request

import (
	"github.com/thedevsaddam/govalidator"
	"gorm.io/datatypes"
)

type LoanData struct {
	Amount float64        `json:"amount"`
	Items  datatypes.JSON `json:"items"`
	TerminalInfo
}

func (l LoanData) Rules() govalidator.MapData {
	return govalidator.MapData{
		"amount": []string{"required", "float"},
		"items":  []string{"required"},
	}
}
