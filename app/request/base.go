package request

import (
	"github.com/thedevsaddam/govalidator"
	"martpay/app/contracts"
)

// Validate - Validate the data structure based on the rules set in its Rules() method.
func Validate(i contracts.ValidationInterface) any {
	opts := govalidator.Options{
		Data:  i,
		Rules: i.Rules(),
	}

	v := govalidator.New(opts)
	if e := v.ValidateStruct(); len(e) > 0 {
		return e
	}

	return nil
}

// TerminalInfo - This data structure refers to the body data that should be present
// in every request that involves creating a transaction.
type TerminalInfo struct {
	Serial  string `json:"serial"`
	VERSION string `json:"VERSION"`
	CHANNEL string `json:"CHANNEL"`
	DEVICE  string `json:"DEVICE"`
	Pin     string `json:"pin"`
}

func (t TerminalInfo) Rules() govalidator.MapData {
	return govalidator.MapData{
		"VERSION": []string{"required"},
		"CHANNEL": []string{"required"},
		"DEVICE":  []string{"required"},
		"pin":     []string{"required", "digits:4", "not_in:0000"},
	}
}
