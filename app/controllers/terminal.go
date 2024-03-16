package ctrl

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"martpay/app/request"
	"martpay/app/utils"
	"martpay/database/query"
	"net/http"
	"strings"
)

func Menus(c echo.Context) error {
	serial := c.Request().Header.Get("deviceId")

	tQ := query.Terminal
	terminal, _ := tQ.Where(tQ.Serial.Eq(serial)).Preload(tQ.Services).First()

	return c.JSON(http.StatusOK, utils.SuccessResponse("Fetched terminal menus", terminal.Menus()))
}

func ResetPin(c echo.Context) error {
	var err error
	pinType := c.Param("type")

	serial := c.Request().Header.Get("deviceId")
	tQ := query.Terminal
	t := tQ.Where(query.Terminal.Serial.Eq(serial))

	if pinType == "pin" { // Reset the terminal pin
		pin := new(request.PinUpdateData)
		if err = c.Bind(pin); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, utils.FailedResponse(err.Error()))
		}

		if vte := request.Validate(pin); vte != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, utils.FailedResponse("Invalid data", vte))
		}

		_, _ = t.Update(tQ.Pin, utils.MakeHash(pin.Pin))
	} else if pinType == "admin-pin" { // Reset the Admin Pin
		pin := new(request.AdminPinUpdateData)
		if err = c.Bind(pin); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, utils.FailedResponse(err.Error()))
		}

		if vte := request.Validate(pin); vte != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, utils.FailedResponse("Invalid data", vte))
		}

		_, _ = t.Update(tQ.AdminPin, utils.MakeHash(pin.AdminPin))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse(fmt.Sprintf("%s Updated.", strings.ToUpper(pinType))))
}
