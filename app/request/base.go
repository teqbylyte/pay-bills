package request

import (
	"github.com/thedevsaddam/govalidator"
	"martpay/app/contracts"
	"martpay/app/utils"
)

// NewTransactionInterface - This defines the methods that every request data which would create a new transaction should implement.
type NewTransactionInterface interface {
	contracts.ValidationInterface

	// GetAmount - The amount for the specific transaction request
	GetAmount() float64

	// GetTerminalInfo - The terminal info in the request
	GetTerminalInfo() TerminalInfo
}

// Validate - Validate the data structure based on the rules set in its Rules() method.
func Validate(i contracts.ValidationInterface) any {
	opts := govalidator.Options{
		Data:  i,
		Rules: i.Rules(),
	}

	v := govalidator.New(opts)
	if e := v.ValidateStruct(); len(e) > 0 {
		return utils.FailedValidationResponse(e)
	}

	return nil
}

func ValidateNewTransaction(i NewTransactionInterface) any {
	var err any

	terminalInfo := i.GetTerminalInfo()
	if err = Validate(&terminalInfo); err != nil {
		return err
	}

	if err = Validate(i); err != nil {
		return err
	}

	return nil
}
