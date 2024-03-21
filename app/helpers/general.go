package helpers

import (
	"github.com/labstack/echo/v4"
	"martpay/app/models"
	"martpay/database/query"
)

// DeviceTerminal - Get the terminal making the request from the deviceId set in the request header.
func DeviceTerminal(c echo.Context) *models.Terminal {
	serial := c.Request().Header.Get("deviceId")
	tQ := query.Terminal
	terminal, _ := tQ.Where(query.Terminal.Serial.Eq(serial)).First()

	return terminal
}
