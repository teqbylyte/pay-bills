package models

type Service struct {
	BaseModel
	ProviderId  uint   `json:"provider_id"`
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	MenuName    string `json:"menu_name"`
	IsAvailable bool   `json:"is_available"`
	Description string `json:"description"`
	Menu        bool   `json:"menu"`
	Internal    bool   `json:"internal"`
}

type ServiceProvider struct {
	BaseModel
	ServiceId uint `json:"service_id" gorm:"not null"`
	Name      uint `json:"name" gorm:"not null"`
}
