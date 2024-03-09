package ctrl

import (
	"github.com/labstack/echo/v4"
	"martpay/app/models"
	"martpay/database"
	"net/http"
)

func Dashboard(c echo.Context) error {
	var transactions []models.Transactions
	db.Db.Find(&transactions)

	//transaction, _ := query.Transactions.First()

	return c.JSON(http.StatusOK, transactions)
}
