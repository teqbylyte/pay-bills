package ctrl

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"martpay/app/models"
	"martpay/database/query"
	"net/http"
)

func Dashboard(c echo.Context) error {
	transactions := make([]models.Transactions, 0)

	err := query.Transactions.
		Scan(&transactions)

	if err != nil {
		fmt.Println("Transaction Error: ", err)
	}

	return c.JSON(http.StatusOK, transactions)
}
