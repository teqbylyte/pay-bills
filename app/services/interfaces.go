package services

import (
	"pay-bills/app/request"
	"pay-bills/app/utils"
)

type ProviderFn func(txRef string, i request.NewTnxInterface) utils.MyResponse

type ProviderInterface interface {
	GetName() string
}

type AirtimeInterface interface {
	ProviderInterface
	PurchaseAirtime(txRef string, i request.NewTnxInterface) utils.MyResponse
}

type DataInterface interface {
	ProviderInterface
	GetDataPlans(network string) utils.MyResponse
	PurchaseData(reference string, data any) utils.MyResponse
}

type TransferInterface interface {
	ProviderInterface
	GetBanks() utils.MyResponse
	ValidateAccount() utils.MyResponse
	MakeTransfer(reference string, data any) utils.MyResponse
}
