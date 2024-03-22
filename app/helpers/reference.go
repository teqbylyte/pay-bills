package helper

import (
	"fmt"
	"martpay/app/utils"
	"martpay/database/query"
	"time"
)

func NewTransactionReference() string {
	reference := fmt.Sprintf("mp_%s", utils.GenerateReference(17))

	// check for duplicate reference
	count, _ := query.Transactions.Where(query.Transactions.Reference.Eq(reference)).Count()
	if count > 0 {
		return NewTransactionReference()
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
