package contracts

import "github.com/thedevsaddam/govalidator"

type ValidationInterface interface {
	// Rules - The validation rules to be used for the body of the request
	Rules() govalidator.MapData
}
