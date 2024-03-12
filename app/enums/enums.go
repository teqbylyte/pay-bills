package enums

type State string
type Status string
type DocumentType string
type ScheduleType string
type SchedulePeriod string

const (
	ACTIVE    State = "ACTIVE"
	INACTIVE  State = "INACTIVE"
	DISABLED  State = "DISABLED"
	SUSPENDED State = "SUSPENDED"

	SUCCESSFUL Status = "SUCCESSFUL"
	PENDING    Status = "PENDING"
	FAILED     Status = "FAILED"

	PassportPhotograph DocumentType = "PASSPORT PHOTOGRAPH"
	IDCard             DocumentType = "ID CARD"
	UtilityBill        DocumentType = "Utility Bill"
)
