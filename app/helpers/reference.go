package helper

import (
	"fmt"
	"pay-bills/app/utils"
	"pay-bills/database/query"
	"time"
)

// NewTnxRef - Unique transaction reference.
func NewTnxRef() string {
	reference := fmt.Sprintf("mp_%s", utils.GenerateReference(17))

	// check for duplicate reference
	count, _ := query.Transactions.Where(query.Transactions.Reference.Eq(reference)).Count()
	if count > 0 {
		return NewTnxRef()
	}

	return reference
}

// NewWalletTnxRef - Unique wallet transaction reference
func NewWalletTnxRef() string {
	reference := fmt.Sprintf("mp_%sw", utils.GenerateReference(16))

	// check for duplicate reference
	count, _ := query.WalletTransaction.Where(query.WalletTransaction.Reference.Eq(reference)).Count()
	if count > 0 {
		return NewTnxRef()
	}

	return reference
}

func NewTid() string {
	year := fmt.Sprintf("%v", time.Now().Year())

	tid := fmt.Sprintf("%s%s", year, utils.GenerateReference(4))

	// check for duplicate tid
	count, _ := query.Terminal.Where(query.Transactions.Reference.Eq(tid)).Count()
	if count > 0 {
		return NewTid()
	}

	return tid
}
