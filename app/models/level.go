package models

type KycLevel struct {
	BaseModel
	Name           string  `json:"name" gorm:"not null"`
	DailyLimit     float64 `json:"daily_limit" gorm:"not null"`
	SingleTransMax float64 `json:"single_trans_max" gorm:"not null"`
	MaxBalance     float64 `json:"max_balance" gorm:"not null"`
}
