package ctrl

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gen"
	"martpay/app/dto"
	"martpay/app/enums"
	"martpay/app/models"
	"martpay/app/utils"
	"martpay/database/query"
	"net/http"
)

func Dashboard(c echo.Context) error {
	user := c.Get("user").(*model.User)

	data := make(map[string]any, 4)

	data["wallet"] = map[string]any{
		"account_number": user.Wallet.AccountNumber,
		"balance":        user.Wallet.Balance,
		"commission":     user.Wallet.Commission,
		"status":         user.Wallet.Status,
		"updated_at":     user.Wallet.UpdatedAt,
	}

	transactions := make([]dto.TransactionDto, 5)

	tQ := query.Transactions                                            // Transactions Query
	agentTq := tQ.Where(tQ.UserId.Eq(user.ID))                          // Agent transactions query
	_ = agentTq.Order(tQ.CreatedAt.Desc()).Limit(5).Scan(&transactions) // Get agent transactions

	data["transactions"] = transactions

	sQ := query.Service                                       // Service query
	successfulTq := agentTq.Where(tQ.Status.Eq("SUCCESSFUL")) // Query for Successful transactions

	cashoutCount, _ := successfulTq.Where(gen.Exists(sQ.Where(sQ.Slug.Eq(enums.CASHOUT)))).Count()

	transferCount, _ := successfulTq.Where(gen.Exists(sQ.Where(sQ.Slug.Eq(enums.TRANSFER)))).Count()

	billPaymentsCount, _ := successfulTq.Where(gen.Exists(sQ.Where(sQ.Slug.Eq(enums.CABLETV)))).
		Or(tQ.Where(gen.Exists(sQ.Where(sQ.Slug.Eq(enums.ELECTRICITY))))).
		Count()

	airtimeCount, _ := successfulTq.Where(gen.Exists(sQ.Where(sQ.Slug.Eq(enums.AIRTIME)))).
		Or(tQ.Where(gen.Exists(sQ.Where(sQ.Slug.Eq(enums.DATA))))).
		Count()

	data["summary"] = map[string]int64{
		"cashout_count":       cashoutCount,
		"transfer_count":      transferCount,
		"bill_payments_count": billPaymentsCount,
		"airtime_count":       airtimeCount,
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Dashboard Loaded", data))
}
