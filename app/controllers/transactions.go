package ctrl

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"martpay/app/models"
	"martpay/app/utils"
	"martpay/database/query"
	"net/http"
)

// Transactions - Get auth user transactions
func Transactions(c echo.Context) error {
	user := c.Get("user").(*models.User)

	transactions := make([]models.Transactions, 0)

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
