package ctrl

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"martpay/app/models"
	"martpay/database/query"
	"net/http"
)

func Dashboard(c echo.Context) error {
	users := make([]models.User, 0)

	err := query.User.
		Scan(&users)

	if err != nil {
		fmt.Println("Transaction Error: ", err)
	}

	return c.JSON(http.StatusOK, users)
}
