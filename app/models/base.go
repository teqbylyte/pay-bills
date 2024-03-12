package models

import "time"

type BaseModel struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BelongsToUser struct {
	UserId uint `json:"user_id" gorm:"not null"`
	User   User `json:"user"`
}
