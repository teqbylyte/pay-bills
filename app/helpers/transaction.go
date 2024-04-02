package helper

import (
	"errors"
	"fmt"
	"net/http"
	"pay-bills/app/enums"
	"pay-bills/app/models"
	"pay-bills/app/request"
	"pay-bills/app/services"
	"pay-bills/app/utils"
	"pay-bills/database/query"
)

func AllowTransaction(terminal *model.Terminal, data request.NewTnxInterface) (int, any) {
	var statusCode int
	var err any

	if terminal.Status != enums.ACTIVE {
		statusCode, err = http.StatusForbidden, errors.New(fmt.Sprintf("Your terminal is %s", terminal.Status))
	}

	if !terminal.User.IsActive() {
		statusCode, err = http.StatusForbidden, errors.New(fmt.Sprintf("Your account is %s", terminal.Status))
	}

	if vte := request.ValidateNewTnx(data); vte != nil {
		statusCode, err = http.StatusUnprocessableEntity, vte
	}

	if e := CheckTerminalPin(terminal, data.GetTerminalInfo().Pin); e != nil {
		statusCode, err = http.StatusForbidden, e
	}

	if statusCode != http.StatusUnprocessableEntity && err != nil {
		err = utils.FailedResponse(err.(error).Error())
	}

	return statusCode, err
}

func GetService(serviceSlug string) (*model.Service, services.ProviderInterface) {
	sQ := query.Service
	service, _ := sQ.Where(sQ.Slug.Eq(serviceSlug)).Preload(sQ.Provider).First()

	var provider services.ProviderInterface

	switch service.Provider.Name {
	case services.Spout{}.GetName():
		provider = new(services.Spout)

	default:
		provider = nil
	}

	return service, provider
}
