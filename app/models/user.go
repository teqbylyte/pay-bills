package model

import (
	"fmt"
	"gorm.io/gorm"
	"pay-bills/app/enums"
	"pay-bills/app/request"
	"pay-bills/app/utils"
	"time"
)

type User struct {
	BaseModel
	LevelId         uint           `json:"level_id"`
	SuperAgentId    *uint          `json:"super_agent_id"`
	FirstName       string         `json:"first_name" gorm:"not null"`
	OtherNames      string         `json:"other_names" gorm:"not null"`
	BusinessName    string         `json:"business_name"`
	Email           string         `json:"email" gorm:"not null"`
	Phone           string         `json:"phone"`
	Type            enums.Role     `json:"type"`
	Gender          string         `json:"gender"`
	Dob             string         `json:"dob"`
	State           string         `json:"state"`
	Country         string         `json:"country"`
	Address         string         `json:"address"`
	Status          enums.State    `json:"status" gorm:"not null;default:ACTIVE"`
	Avatar          string         `json:"avatar"`
	Bvn             string         `json:"bvn"`
	Nin             string         `json:"nin"`
	ReferralCode    string         `json:"referral_code"`
	EmailVerifiedAt time.Time      `json:"email_verified_at"`
	Password        string         `json:"password" gorm:"not null"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
	Wallet          Wallet         `json:"wallet"`
}

func (u User) Name() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.OtherNames)
}

func (u User) IsActive() bool {
	return u.Status == enums.ACTIVE
}

// NewUser - Create a new user from the request data.
func NewUser(data *request.RegisterData, superAgentId *uint) *User {
	return &User{
		SuperAgentId: superAgentId,
		FirstName:    data.FirstName,
		OtherNames:   data.OtherNames,
		BusinessName: data.BusinessName,
		Phone:        data.Phone,
		Email:        data.Email,
		Dob:          data.Dob,
		State:        data.State,
		Address:      data.Address,
		Gender:       data.Gender,
		Bvn:          data.Bvn,
		Type:         data.Role,
		Password:     utils.MakeHash(data.Password),
	}
}
