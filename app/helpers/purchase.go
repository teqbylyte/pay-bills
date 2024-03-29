package helper

import (
	"martpay/app/models"
	"martpay/app/request"
	"martpay/app/services"
	"martpay/app/utils"
	"net/http"
)

// ProcessPurchase - Process the purchase for the specified service with the provider call.
func ProcessPurchase(
	terminal *models.Terminal,
	service *models.Service,
	tnx *models.Transactions,
	data request.NewTnxInterface,
	provider services.ProviderFn,
) (int, utils.MyResponse) {
	var res utils.MyResponse
	var err error

	commission := 0.0 //  Todo: Get the commission
	// Debit the wallet
	err = ProcessWalletDebit(
		&terminal.User.Wallet,
		service,
		tnx.Amount,
		tnx.Charge,
		tnx.Reference,
		tnx.Info,
		tnx.Amount-commission,
		false,
	)

	if err != nil {
		return http.StatusBadRequest, utils.FailedResponse(err.Error())
	}

	res = provider(tnx.Reference, data)

	return http.StatusOK, res
}
