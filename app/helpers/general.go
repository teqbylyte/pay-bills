package helper

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"pay-bills/app/enums"
	"pay-bills/app/models"
	"pay-bills/app/utils"
	"pay-bills/database/query"
)

// AuthTerminal - Get the terminal making the request from the deviceId set in the request header.
func AuthTerminal(c echo.Context) *model.Terminal {
	serial := c.Request().Header.Get("deviceId")
	tQ := query.Terminal
	terminal, _ := tQ.Where(query.Terminal.Serial.Eq(serial)).Preload(tQ.User.Wallet).First()

	return terminal
}

func CheckTerminalPin(terminal *model.Terminal, pin string) error {
	if terminal.Pin == "0000" {
		if pin != terminal.Pin {
			return errors.New("Incorrect pin.")
		}
	}
	tQ := query.Terminal

	if !utils.HashIsValid(pin, terminal.Pin) {

		// Increase wrong pin count if the wrong pin was entered.
		wrongPinCount := terminal.WrongPinCount + 1
		_, _ = tQ.Where(query.Terminal.Serial.Eq(terminal.Serial)).Update(tQ.WrongPinCount, wrongPinCount)

		// Deactivate the terminal if it exceeds 4 wrong pins.
		if wrongPinCount >= 4 {
			_, _ = tQ.Where(query.Terminal.Serial.Eq(terminal.Serial)).Update(tQ.Status, enums.INACTIVE)
			return errors.New("You entered the wrong pin 4 times. This terminal has been deactivated.")
		}

		return errors.New(fmt.Sprintf("Incorrect Pin. You have %v more trials", 4-wrongPinCount))
	}

	if terminal.WrongPinCount > 0 {
		_, _ = tQ.Where(query.Terminal.Serial.Eq(terminal.Serial)).Update(tQ.WrongPinCount, 0)
	}

	return nil
}

func CheckTerminalAdminPin(terminal *model.Terminal, pin string) error {
	if terminal.Pin == "0000" {
		if pin != terminal.Pin {
			return errors.New("Incorrect admin pin.")
		}
	}

	if !utils.HashIsValid(pin, terminal.AdminPin) {
		return errors.New(fmt.Sprintf("Incorrect admin pin. Try again."))
	}

	return nil
}
