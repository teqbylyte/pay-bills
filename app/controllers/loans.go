package ctrl

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"martpay/app/dto"
	"martpay/app/models"
	"martpay/app/utils"
	"martpay/database/query"
	"net/http"
)

func GetLoans(c echo.Context) error {
	user := c.Get("user").(*models.User)

	var loans []dto.LoanDto

	lQ := query.Loan
	_ = lQ.Where(lQ.UserId.Eq(user.ID)).Scan(&loans)

	fmt.Println(user.ID)
	return c.JSON(http.StatusOK, utils.SuccessResponse("Fetched loan requests", loans))
}
