package models

import (
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"martpay/app/enums"
	"time"
)

type Terminal struct {
	BaseModel
	UserId        uint        `json:"user_id" gorm:"not null"`
	GroupId       uint        `json:"group_id" gorm:"not null"`
	Device        string      `json:"device" gorm:"not null"`
	Serial        string      `json:"serial" gorm:"not null"`
	Status        enums.State `json:"status" gorm:"not null"`
	Tid           string      `json:"tid"`
	Mid           string      `json:"mid"`
	Tmk           string      `json:"tmk"`
	Tsk           string      `json:"tsk"`
	Tpk           string      `json:"tpk"`
	DateTime      string      `json:"date_time"`
	Timeout       int         `json:"timeout"`
	CurrencyCode  string      `json:"currency_code"`
	CountryCode   string      `json:"country_code"`
	CategoryCode  string      `json:"category_code"`
	NameLocation  string      `json:"name_location"`
	AdminPin      string      `json:"admin_pin"`
	Pin           string      `json:"pin"`
	WrongPinCount int         `json:"wrong_pin_count"`
	HasChangedPin bool        `json:"has_changed_pin"`
	Services      []Service   `json:"services" gorm:"many2many:service_terminal"`
	User          User        `json:"agent"`
}

// GenerateToken - Generate valid jwt to be used by the user of the terminal
func (t Terminal) GenerateToken() (string, time.Time) {
	expiresAt := time.Now().Add(time.Hour * 168)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     t.User.ID,
		"serial": t.Serial,
		"exp":    expiresAt.Unix(),
	})

	accessToken, err := token.SignedString([]byte(viper.GetString("APP_KEY")))

	if err != nil {
		panic(err)
	}

	return accessToken, expiresAt
}

// Menus - Get the terminal menus from the services added to the terminal
func (t Terminal) Menus() []map[string]any {
	menus := make([]map[string]any, len(t.Services))

	for i, service := range t.Services {
		menus[i] = map[string]any{
			"id":   service.ID,
			"name": service.MenuName,
		}
	}

	return menus
}
