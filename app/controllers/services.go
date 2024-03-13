package ctrl

import (
	"github.com/labstack/echo/v4"
	"martpay/app/utils"
	"martpay/database/query"
	"net/http"
)

// Services - Get auth user services available to terminal device
func Services(c echo.Context) error {
	serial := c.Request().Header.Get("deviceId")

	tQ := query.Terminal
	terminal, _ := tQ.Where(tQ.Serial.Eq(serial)).Preload(tQ.Services).First()

	return c.JSON(http.StatusOK, utils.SuccessResponse("Fetched terminal menus", terminal.Menus()))
}
