package request

import (
	"github.com/thedevsaddam/govalidator"
	"martpay/app/contracts"
)

// Validate - Validate the struct argument based on the given rules set.
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
