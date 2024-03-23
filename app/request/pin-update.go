package request

import (
	"fmt"
	"github.com/thedevsaddam/govalidator"
)

type PinUpdateData struct {
	Pin        string `json:"pin"`
	ConfirmPin string `json:"confirm_pin"`
	CurrentPin string `json:"current_pin"`
}

func (p PinUpdateData) Rules() govalidator.MapData {
	return govalidator.MapData{
		"pin":         []string{"required", "digits:4"},
		"confirm_pin": []string{"required", fmt.Sprintf("confirm:%s", p.Pin)},
		"current_pin": []string{"required", "digits:4"}, // TODO: Validate that it matches with the pin in the db
	}
}

type AdminPinUpdateData struct {
	AdminPin        string `json:"admin_pin"`
	ConfirmAdminPin string `json:"confirm_admin_pin"`
	CurrentAdminPin string `json:"current_admin_pin"`
}

func (p AdminPinUpdateData) Rules() govalidator.MapData {
	return govalidator.MapData{
		"admin_pin":         []string{"required", "digits:4"},
		"confirm_admin_pin": []string{"required", fmt.Sprintf("confirm:%s", p.AdminPin)},
		"current_admin_pin": []string{"required", "digits:4"}, // TODO: Validate that it matches with the admin_pin in the db
	}
}
