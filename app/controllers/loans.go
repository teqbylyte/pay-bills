package ctrl

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pay-bills/app/dto"
	"pay-bills/app/enums"
	"pay-bills/app/helpers"
	"pay-bills/app/models"
	"pay-bills/app/request"
	"pay-bills/app/utils"
	"pay-bills/database"
	"pay-bills/database/query"
)

func GetLoans(c echo.Context) error {
	user := c.Get("user").(*model.User)

	var loans []dto.LoanDto

	lQ := query.Loan
	_ = lQ.Where(lQ.UserId.Eq(user.ID)).Scan(&loans)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Fetched loan requests", loans))
}

func CreateLoan(c echo.Context) error {
	agent := c.Get("user").(*model.User)
	terminal := helper.AuthTerminal(c)

	data := new(request.LoanData)
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.FailedResponse(err.Error()))
	}

	if vte := request.ValidateNewTnx(data); vte != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, vte)
	}

	if err := helper.CheckTerminalPin(terminal, data.TerminalInfo.Pin); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.FailedResponse(err.Error()))
	}

	service, _ := query.Service.Where(query.Service.Slug.Eq(enums.LOAN)).First()

	charge := 0.0 // Todo: Get Charge
	totalAmount := charge + data.Amount
	reference := helper.NewTnxRef()
	info := "Loan request"

	transaction := model.NewPendingTransaction(terminal, service, data.Amount, charge, reference, info,
		"INTERNAL", "", &data.TerminalInfo)

	// TODO: Check if loan can be created. (1) Has super agent (2) Does not have pending loan

	dbTxErr := query.Use(database.Db).Transaction(func(tx *query.Query) error {
		if err := tx.Transactions.Create(transaction); err != nil {
			return err
		}

		loan := model.Loan{
			UserId:        agent.ID,
			TransactionId: transaction.ID,
			Amount:        data.Amount,
			Charge:        0,
			AmountOwed:    totalAmount,
			Items:         data.Items,
			Status:        enums.PENDING,
		}

		if err := tx.Loan.Create(&loan); err != nil {
			return err
		}

		return nil
	})

	if dbTxErr != nil {
		return c.JSON(http.StatusBadRequest, utils.FailedResponse(dbTxErr.Error()))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Loan request created..."))
}
