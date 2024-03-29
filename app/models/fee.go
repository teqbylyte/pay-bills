package model

import "gorm.io/datatypes"

type Fee struct {
	BaseModel
	GroupId    uint           `json:"group_id" gorm:"not null"`
	ServiceId  string         `json:"service_id" gorm:"not null"`
	Title      string         `json:"title" gorm:"not null"`
	Type       string         `json:"type" gorm:"not null"`
	Amount     float64        `json:"amount" gorm:"not null"`
	AmountType string         `json:"amount_type" gorm:"not null"`
	Cap        float64        `json:"cap" gorm:"default:0.0"`
	Info       string         `json:"info"`
	Config     datatypes.JSON `json:"config"`
	Structure  datatypes.JSON `json:"structure"`
}
