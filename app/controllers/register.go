package ctrl

import (
	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	helper "martpay/app/helpers"
	"martpay/app/models"
	"martpay/app/request"
	"martpay/app/utils"
	"martpay/database"
	"martpay/database/query"
	"net/http"
)

func Register(c echo.Context) error {
	var data request.RegisterData
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailedResponse("An error occurred", err.Error()))
	}

	if vte := request.Validate(&data); vte != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, utils.FailedValidation(vte))
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

	agent := models.NewUser(&data, superAgentId)

	dbTxErr := query.Use(database.Db).Transaction(func(tx *query.Query) error {
		if err := tx.User.Create(agent); err != nil {
			return err
		}

		uwid, _ := uuid.NewV4()

		wallet := models.Wallet{
			UserId:        agent.ID,
			Uwid:          uwid.String(),
			AccountNumber: data.Phone[len(data.Phone)-10:],
		}

		if err := tx.Wallet.Create(&wallet); err != nil {
			return err
		}

		// TODO: Create Virtual Account

		terminal := models.NewTerminal(agent, data.Serial, data.Device, helper.NewTid())

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
