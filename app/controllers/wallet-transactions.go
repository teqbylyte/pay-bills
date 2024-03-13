package ctrl

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"martpay/app/models"
	"martpay/app/utils"
	"martpay/database/query"
	"net/http"
)

func WalletTransactions(c echo.Context) error {
	user := c.Get("user").(*models.User)

	transactions := make([]models.WalletTransaction, 0)

	wtQ := query.WalletTransaction
	err := wtQ.Where(wtQ.WalletId.Eq(user.Wallet.ID)).
		Order(wtQ.CreatedAt.Desc()).
		Scan(&transactions)

	if err != nil {
		fmt.Println("Transaction Error: ", err)
		return c.JSON(http.StatusBadRequest, utils.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Wallet Transactions fetched", transactions))
}
