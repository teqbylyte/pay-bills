package contracts

import "github.com/thedevsaddam/govalidator"

type ValidationInterface interface {
	Rules() govalidator.MapData
}
