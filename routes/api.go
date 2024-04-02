package routes

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	ctrl "pay-bills/app/controllers"
	serviceCtrl "pay-bills/app/controllers/services"
	"pay-bills/database/query"
)

func Api(e *echo.Echo) {
	// Create a route prefix v1 - version 1.
	v1 := e.Group("/v1")
	v1.Use(middleware.RequestID())
	v1.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST},
	}))

	//Register route
	v1.POST("/register", ctrl.Register)

	// Login route
	v1.POST("/auth", ctrl.Login)

	// Authenticated routes
	auth := v1.Group("", echojwt.WithConfig(jwtConfig()))

	auth.GET("/dashboard", ctrl.Dashboard)
	auth.GET("/transactions", ctrl.Transactions)
	auth.GET("/wallet-transactions", ctrl.WalletTransactions)
	auth.GET("/terminal/menus", ctrl.Menus)
	auth.POST("/terminal/reset/:type", ctrl.ResetPin)
	auth.GET("/loans", ctrl.GetLoans)

	// Request for terminal transactions and purchases
	//TODO: Check that the serial in the deviceId belongs to the auth user

	auth.POST("/loans", ctrl.CreateLoan)

	service := auth.Group("/services")

	service.GET("/airtime/networks", serviceCtrl.GetAirtimeNetworks)
	service.POST("/airtime/purchase", serviceCtrl.PurchaseAirtime)
}

func jwtConfig() echojwt.Config {
	return echojwt.Config{
		SigningKey: []byte(viper.GetString("APP_KEY")),
		SuccessHandler: func(c echo.Context) {
			claims := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)
			userId := uint(claims["id"].(float64)) // Get user id from claims

			users := query.User
			user, err := users.Where(users.ID.Eq(userId)).Preload(users.Wallet).First()

			if err == nil {
				c.Set("user", user) // Set auth user to the request context
			}
		},
	}
}
