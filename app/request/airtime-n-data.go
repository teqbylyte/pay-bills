package request

import (
	"github.com/thedevsaddam/govalidator"
)

type AirtimeData struct {
	Network string  `json:"network" xml:"network"`
	Amount  float64 `json:"amount" xml:"amount"`
	Phone   string  `json:"phone" xml:"phone"`
	TerminalInfo
}

func (a AirtimeData) Rules() govalidator.MapData {
	return govalidator.MapData{
		"amount":  []string{"required", "float", "min:50"},
		"network": []string{"required", "in:mtn,airtel,9mobile,glo"},
		"phone":   []string{"required", "digits:11"},
	}
}

func (a AirtimeData) GetAmount() float64 {
	return a.Amount
}

func (a AirtimeData) GetTerminalInfo() TerminalInfo {
	return a.TerminalInfo
}
