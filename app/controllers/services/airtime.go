package serviceCtrl

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"pay-bills/app/helpers"
	"pay-bills/app/models"
	"pay-bills/app/request"
	"pay-bills/app/services"
	"pay-bills/app/utils"
	"pay-bills/database/query"
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

	reference := helper.NewTnxRef()
	info := fmt.Sprintf("Airtime purchase of %v for %s", data.Amount, data.Network)
	charge := 0.0 // TODO: Get charge from terminal group

	transaction := model.NewPendingTransaction(terminal, service, data.Amount, data.Amount+charge, reference, info,
		provider.GetName(), data.Phone, &data.TerminalInfo)
	// Record transaction in the db.
	_ = query.Transactions.Create(transaction)

	airtimeProvider := provider.(services.AirtimeInterface)

	status, response := helper.ProcessPurchase(terminal, service, transaction, data,
		airtimeProvider.PurchaseAirtime,
	)

	return c.JSON(status, response)
}
