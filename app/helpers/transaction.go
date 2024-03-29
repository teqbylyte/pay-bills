package helper

import (
	"errors"
	"fmt"
	"martpay/app/contracts"
	"martpay/app/enums"
	"martpay/app/models"
	"martpay/app/request"
	"martpay/app/services"
	"martpay/app/utils"
	"martpay/database/query"
	"net/http"
)

func AllowTransaction(terminal *models.Terminal, i request.NewTransactionInterface) (int, any) {
	var statusCode int
	var err any

	if terminal.Status != enums.ACTIVE {
		statusCode, err = http.StatusForbidden, errors.New(fmt.Sprintf("Your terminal is %s", terminal.Status))
	}

	if !terminal.User.IsActive() {
		statusCode, err = http.StatusForbidden, errors.New(fmt.Sprintf("Your account is %s", terminal.Status))
	}

	if vte := request.ValidateNewTransaction(i); vte != nil {
		statusCode, err = http.StatusUnprocessableEntity, vte
	}

	if e := CheckTerminalPin(terminal, i.GetTerminalInfo().Pin); e != nil {
		statusCode, err = http.StatusForbidden, e
	}

	if statusCode != http.StatusUnprocessableEntity && err != nil {
		err = utils.FailedResponse(err.(error).Error())
	}

	return statusCode, err
}

func GetService(serviceSlug string) (*models.Service, contracts.BaseProviderInterface) {
	sQ := query.Service
	service, _ := sQ.Where(sQ.Slug.Eq("airtime")).Preload(sQ.Provider).First()

	var provider contracts.BaseProviderInterface

	switch service.Provider.Name {
	case services.Spout{}.GetName():
		provider = new(services.Spout)

	default:
		provider = nil
	}

	return service, provider
}
