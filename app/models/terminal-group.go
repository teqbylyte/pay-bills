package models

type TerminalGroup struct {
	BaseModel
	Name string `json:"name" gorm:"not null"`
	Info string `json:"info"`
}
