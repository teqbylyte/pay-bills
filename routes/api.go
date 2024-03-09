package routes

import (
	"github.com/labstack/echo/v4"
	ctrl "martpay/app/controllers"
)

func Api(e *echo.Echo) {
	// Create a route prefix v1 - version 1.
	v1 := e.Group("/v1")

	v1.GET("/dashboard", ctrl.Dashboard)
}
