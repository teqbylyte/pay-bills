package contracts

import "martpay/app/utils"

type BaseProviderInterface interface {
	GetName() string
}

type AirtimeInterface interface {
	BaseProviderInterface
	PurchaseAirtime(amount string, network string, phone string, reference string) utils.MyResponse
}

type DataInterface interface {
	BaseProviderInterface
	GetDataPlans(network string) utils.MyResponse
	PurchaseData(amount string, network string, phone string, planCode string, reference string, meta ...any) utils.MyResponse
}

type TransferInterface interface {
	BaseProviderInterface
	GetBanks() utils.MyResponse
	ValidateAccount() utils.MyResponse
	MakeTransfer(amount string, network string, phone string, reference string) utils.MyResponse
}
