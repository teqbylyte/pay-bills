package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	BaseModel
	LevelId         uint           `json:"level_id"`
	SuperAgentId    uint           `json:"super_agent_id"`
	FirstName       string         `json:"first_name" gorm:"not null"`
	OtherNames      string         `json:"other_names" gorm:"not null"`
	BusinessName    string         `json:"business_name"`
	Email           string         `json:"email" gorm:"not null"`
	Phone           string         `json:"phone"`
	Gender          string         `json:"gender"`
	Dob             string         `json:"dob"`
	State           string         `json:"state"`
	Country         string         `json:"country"`
	Address         string         `json:"address"`
	Status          string         `json:"status" gorm:"not null"`
	Avatar          string         `json:"avatar"`
	Bvn             string         `json:"bvn"`
	Nin             string         `json:"nin"`
	ReferralCode    string         `json:"referral_code"`
	EmailVerifiedAt time.Time      `json:"email_verified_at"`
	Password        string         `json:"password" gorm:"not null"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
}
