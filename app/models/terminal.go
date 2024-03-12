package models

type Terminal struct {
	BaseModel
	BelongsToUser
	GroupId       uint   `json:"group_id" gorm:"not null"`
	Device        string `json:"device" gorm:"not null"`
	Serial        string `json:"serial" gorm:"not null"`
	Status        string `json:"status" gorm:"not null"`
	Tid           string `json:"tid"`
	Mid           string `json:"mid"`
	Tmk           string `json:"tmk"`
	Tsk           string `json:"tsk"`
	Tpk           string `json:"tpk"`
	DateTime      string `json:"date_time"`
	Timeout       int    `json:"timeout"`
	CurrencyCode  string `json:"currency_code"`
	CountryCode   string `json:"country_code"`
	CategoryCode  string `json:"category_code"`
	NameLocation  string `json:"name_location"`
	AdminPin      string `json:"admin_pin"`
	Pin           string `json:"pin"`
	WrongPinCount int    `json:"wrong_pin_count"`
	HasChangedPin bool   `json:"has_changed_pin"`
}
