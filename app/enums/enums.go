package enums

type State string
type Status string
type DocumentType string
type ScheduleType string
type SchedulePeriod string
type Role string
type TransType string
type ACTION string

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

	CASHOUT        string = "cashoutwithdrawal"
	AIRTIME        string = "airtime"
	DATA           string = "internetdata"
	TRANSFER       string = "banktransfer"
	WALLETTRANSFER string = "wallettransfer"
	LOAN           string = "loan"
	ELECTRICITY    string = "electricity"
	CABLETV        string = "cabeltv"
	FUNDING        string = "fundinginbound"

	AGENT      Role = "agent"
	SUPERAGENT Role = "super-agent"

	CHARGE      TransType = "CHARGE"
	COMMISSION  TransType = "COMMISSION"
	TRANSACTION TransType = "TRANSACTION"

	CREDIT ACTION = "CREDIT"
	DEBIT  ACTION = "DEBIT"
)
