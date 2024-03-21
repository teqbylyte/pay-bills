package ctrl

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	state "martpay/app/enums"
	"martpay/app/models"
	"martpay/app/utils"
	"martpay/database/query"
	"net/http"
)

type LoginData struct {
	Serial   string `json:"serial"`
	Password string `json:"password"`
}

func Login(c echo.Context) error {
	var loginData LoginData
	err := c.Bind(&loginData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailedResponse("Request error", err.Error()))
	}

	terminals := query.Terminal
	terminal, err := terminals.Where(terminals.Serial.Eq(loginData.Serial)).Preload(terminals.User.Wallet).First()
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.FailedResponse("Terminal not found."))
	}

	agent := terminal.User // Get the terminal agent

	if !utils.HashIsValid(loginData.Password, agent.Password) {
		return c.JSON(http.StatusUnprocessableEntity, utils.FailedResponse("Invalid credentials."))
	}

	err = ensureTerminalCanLogin(terminal)

	if err != nil {
		return c.JSON(http.StatusForbidden, utils.FailedResponse(err.Error()))
	}

	token, expiresAt := terminal.GenerateToken()

	return c.JSON(http.StatusOK, utils.SuccessResponse("Login successful", map[string]any{
		"id":              terminal.ID,
		"serial":          terminal.Serial,
		"tid":             terminal.Tid,
		"mid":             terminal.Mid,
		"tmk":             terminal.Tmk,
		"tpk":             terminal.Tpk,
		"tsk":             terminal.Tsk,
		"country_code":    terminal.CountryCode,
		"currency_code":   terminal.CurrencyCode,
		"first_name":      agent.FirstName,
		"name":            agent.Name(),
		"business_name":   agent.BusinessName,
		"phone":           agent.Phone,
		"level":           agent.LevelId,
		"terminal_status": terminal.Status,
		"access_token": map[string]any{
			"value":      token,
			"type":       "Bearer",
			"expires_at": expiresAt,
		},
	}))
}

// Check that the terminal status, the terminal agent status
// and the agent wallet status are all active.
func ensureTerminalCanLogin(terminal *models.Terminal) error {
	if terminal.Status != state.ACTIVE {
		return errors.New(fmt.Sprintf("Your terminal is currently %s", terminal.Status))
	}

	if terminal.User.Status != state.ACTIVE {
		return errors.New(fmt.Sprintf("Your account is currently %s", terminal.User.Status))
	}

	if terminal.User.Wallet.Status != state.ACTIVE {
		return errors.New(fmt.Sprintf("Your wallet is currently %s", terminal.User.Wallet.Status))
	}

	return nil
}
