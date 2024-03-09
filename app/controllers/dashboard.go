package ctrl

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Dashboard(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Mart Pay dashboard!")
}
