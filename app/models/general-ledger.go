package models

type GeneralLedger struct {
	BaseModel
	ServiceId uint    `json:"service_id" gorm:"not null"`
	Balance   float64 `json:"balance" gorm:"default:0.0"`
}

type GlTransaction struct {
	FromUserId  uint    `json:"from_user_id"`
	GlId        uint    `json:"gl_id" gorm:"not null"`
	Amount      float64 `json:"amount" gorm:"not null"`
	PrevBalance float64 `json:"prev_balance" gorm:"not null"`
	NewBalance  float64 `json:"new_balance" gorm:"not null"`
	Type        string  `json:"type" gorm:"not null"`
	Info        string  `json:"info"`
}
