package models

import "gorm.io/datatypes"

type Loan struct {
	BaseModel
	BelongsToUser
	TransactionId uint           `json:"transaction_id" gorm:"not null"`
	Amount        float64        `json:"amount" gorm:"not null"`
	Charge        float64        `json:"charge" gorm:"default:0.0"`
	AmountOwed    float64        `json:"amount_owed" gorm:"default:0.0"`
	Items         datatypes.JSON `json:"items" gorm:"not null"`
	Status        string         `json:"status" gorm:"not null"`
	Info          *string        `json:"info"`
	DeclineReason *string        `json:"decline_reason"`
	ApprovedBy    *uint          `json:"approved_by"`
	ConfirmedBy   *uint          `json:"confirmed_by"`
	DeclinedBy    *uint          `json:"declined_by"`
}
