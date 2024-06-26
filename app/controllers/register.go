package ctrl

import (
	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	helper "pay-bills/app/helpers"
	"pay-bills/app/models"
	"pay-bills/app/request"
	"pay-bills/app/utils"
	"pay-bills/database"
	"pay-bills/database/query"
)

func Register(c echo.Context) error {
	var data request.RegisterData
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailedResponse("An error occurred", err.Error()))
	}

	if vte := request.Validate(&data); vte != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, vte)
	}

	uQ := query.User

	var superAgentId *uint

	// Get the superagent from the referral code passed.
	if len(data.ReferralCode) > 0 {
		superAgent, err := uQ.Where(uQ.ReferralCode.Eq(data.ReferralCode)).First()
		if err != nil {
			superAgentId = &superAgent.ID
		}
	}

	agent := model.NewUser(&data, superAgentId)

	dbTxErr := query.Use(database.Db).Transaction(func(tx *query.Query) error {
		if err := tx.User.Create(agent); err != nil {
			return err
		}

		uwid, _ := uuid.NewV4()

		wallet := model.Wallet{
			UserId:        agent.ID,
			Uwid:          uwid.String(),
			AccountNumber: data.Phone[len(data.Phone)-10:],
		}

		if err := tx.Wallet.Create(&wallet); err != nil {
			return err
		}

		// TODO: Create Virtual Account

		terminal := model.NewTerminal(agent, data.Serial, data.Device, helper.NewTid())

		if err := tx.Terminal.Create(terminal); err != nil {
			return err
		}

		return nil
	})

	if dbTxErr != nil {
		return c.JSON(http.StatusBadRequest, utils.FailedResponse("An error occurred during registration", dbTxErr.Error()))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Registration successful!"))
}
