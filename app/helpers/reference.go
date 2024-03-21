package helpers

import (
	"fmt"
	"martpay/database/query"
	"math/rand"
	"time"
	"unsafe"
)

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	letterBytes   = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func GenerateReference(size int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, size)

	for i, cache, remain := size-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

func NewTransactionReference() string {
	reference := fmt.Sprintf("mp_%s", GenerateReference(17))

	// check for duplicate reference
	count, _ := query.Transactions.Where(query.Transactions.Reference.Eq(reference)).Count()
	if count > 0 {
		return NewTransactionReference()
	}

	return reference
}
