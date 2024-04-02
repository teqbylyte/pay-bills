package request

import (
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"pay-bills/app/enums"
	"strings"
)

// RegisterData - The data structure accepted during registration of new agent.
type RegisterData struct {
	Serial          string     `json:"serial"`
	Device          string     `json:"device"`
	FirstName       string     `json:"first_name"`
	OtherNames      string     `json:"other_names"`
	BusinessName    string     `json:"business_name"`
	Email           string     `json:"email"`
	Phone           string     `json:"phone"`
	Gender          string     `json:"gender"`
	Dob             string     `json:"dob"`
	State           string     `json:"state"`
	Role            enums.Role `json:"role"`
	Address         string     `json:"address"`
	Password        string     `json:"password"`
	ConfirmPassword string     `json:"confirm_password"`
	Bvn             string     `json:"bvn"`
	ReferralCode    string     `json:"referral_code"`
}

func (r RegisterData) Rules() govalidator.MapData {
	states := strings.Join(enums.States(), ",")

	return govalidator.MapData{
		"serial":           []string{"required", "unique:terminals"},
		"device":           []string{"required"},
		"first_name":       []string{"required"},
		"other_names":      []string{"required"},
		"email":            []string{"required", "email", "unique:users"},
		"phone":            []string{"required", "digits_between:11,15", "unique:users"},
		"gender":           []string{"required", "in:MALE,FEMALE"},
		"dob":              []string{"required", "date"},
		"state":            []string{"required", fmt.Sprintf("in:%s", states)},
		"role":             []string{"required", "in:agent,super-agent"},
		"address":          []string{"required"},
		"password":         []string{"required"},
		"confirm_password": []string{"required", fmt.Sprintf("confirm:%s", r.Password)}, // TODO: confirm password
		"bvn":              []string{"required", "digits:11"},
		"business_name":    []string{"required"},
		"referral_code":    []string{"exists:users"},
	}
}
