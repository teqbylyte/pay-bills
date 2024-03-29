package model

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"martpay/app/enums"
	"strings"
	"time"
)

type Terminal struct {
	BaseModel
	UserId        uint        `json:"user_id" gorm:"not null"`
	GroupId       uint        `json:"group_id" gorm:"not null"`
	Device        string      `json:"device" gorm:"not null"`
	Serial        string      `json:"serial" gorm:"not null"`
	Status        enums.State `json:"status" gorm:"not null;default:ACTIVE"`
	Tid           string      `json:"tid"`
	Mid           string      `json:"mid"`
	Tmk           string      `json:"tmk"`
	Tsk           string      `json:"tsk"`
	Tpk           string      `json:"tpk"`
	DateTime      string      `json:"date_time"`
	Timeout       int         `json:"timeout"`
	CurrencyCode  string      `json:"currency_code" gorm:"default:566"`
	CountryCode   string      `json:"country_code" gorm:"default:566"`
	CategoryCode  string      `json:"category_code" gorm:"default:1234"`
	NameLocation  string      `json:"name_location"`
	AdminPin      string      `json:"admin_pin" gorm:"default:0000"`
	Pin           string      `json:"pin" gorm:"default:0000"`
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

// NewTerminal - Creating a new terminal
func NewTerminal(agent *User, serial string, device string, tid string) *Terminal {
	year := fmt.Sprintf("%v", time.Now().Year())
	mid := year + strings.Repeat("0", 11)

	nameLocation := fmt.Sprintf("%s            LA NG", agent.Name())

	return &Terminal{
		UserId:       agent.ID,
		Serial:       serial,
		Device:       device,
		Tid:          tid,
		Mid:          mid,
		NameLocation: nameLocation,
	}
}
