package serviceCtrl

import (
	"fmt"
	"github.com/labstack/echo/v4"
	helper "martpay/app/helpers"
	"martpay/app/models"
	"martpay/app/request"
	"martpay/app/utils"
	"martpay/database/query"
	"net/http"
)

func GetAirtimeNetworks(c echo.Context) error {
	data := []map[string]string{
		{"id": "mtn", "name": "MTN"},
		{"id": "airtel", "name": "AIRTEL"},
		{"id": "9mobile", "name": "9MOBILE"},
		{"id": "glo", "name": "GLO"},
	}
	return c.JSON(http.StatusOK, utils.SuccessResponse("Fetched airtime networks", data))
}

func PurchaseAirtime(c echo.Context) error {
	terminal := helper.AuthTerminal(c)

	var data request.AirtimeData
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.FailedResponse(err.Error()))
	}

	if statusCode, err := helper.AllowTransaction(terminal, &data); err != nil {
		return echo.NewHTTPError(statusCode, err)
	}

	service, provider := helper.GetService("airtime")
	if provider == nil || service == nil {
		return echo.NewHTTPError(http.StatusNotFound, utils.FailedResponse("Service Unavailable"))
	}

	reference := helper.NewTransactionReference()
	info := fmt.Sprintf("Airtime purchase of %v for %s", data.Amount, data.Network)

	transaction := models.NewPendingTransaction(terminal, service, data.Amount, data.Amount, reference, info,
		provider.GetName(), data.Phone, &data.TerminalInfo)
	// Record in the db.
	_ = query.Transactions.Create(transaction)

	// Todo: Debit the wallet

	return c.JSON(http.StatusOK, utils.SuccessResponse("Airtime purchase successful"))
}
