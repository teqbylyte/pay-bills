package ctrl

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"pay-bills/app/models"
	"pay-bills/app/utils"
	"pay-bills/database/query"
)

// WalletTransactions - Get auth user wallet transactions
func WalletTransactions(c echo.Context) error {
	user := c.Get("user").(*model.User)

	transactions := make([]model.WalletTransaction, 0)

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
