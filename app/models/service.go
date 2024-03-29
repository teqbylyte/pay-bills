package model

type Service struct {
	BaseModel
	ProviderId  uint            `json:"provider_id"`
	Slug        string          `json:"slug"`
	Name        string          `json:"name"`
	MenuName    string          `json:"menu_name"`
	IsAvailable bool            `json:"is_available"`
	Description string          `json:"description"`
	Menu        bool            `json:"menu"`
	Internal    bool            `json:"internal"`
	Provider    ServiceProvider `json:"provider" gorm:"foreignKey:ProviderId"`
}

type ServiceProvider struct {
	BaseModel
	ServiceId uint   `json:"service_id" gorm:"not null"`
	Name      string `json:"name" gorm:"not null"`
}
