package ctrl

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"pay-bills/app/models"
	"pay-bills/app/utils"
	"pay-bills/database/query"
)

// Transactions - Get auth user transactions
func Transactions(c echo.Context) error {
	user := c.Get("user").(*model.User)

	transactions := make([]model.Transactions, 0)

	tQ := query.Transactions
	err := tQ.Where(tQ.UserId.Eq(user.ID)).
		Order(tQ.CreatedAt.Desc()).
		Scan(&transactions)

	if err != nil {
		fmt.Println("Transaction Error: ", err)
		return c.JSON(http.StatusBadRequest, utils.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Transactions fetched", transactions))
}
