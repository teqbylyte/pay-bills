package routes

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	ctrl "martpay/app/controllers"
	"martpay/database/query"
)

func Api(e *echo.Echo) {
	// Create a route prefix v1 - version 1.
	v1 := e.Group("/v1")
	v1.Use(middleware.RequestID())
	v1.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST},
	}))

	// Login route
	v1.POST("/auth", ctrl.Login)

	// Authenticated routes
	auth := v1.Group("", echojwt.WithConfig(jwtConfig()))

	auth.GET("/dashboard", ctrl.Dashboard)
}

func jwtConfig() echojwt.Config {
	return echojwt.Config{
		SigningKey: []byte(viper.GetString("APP_KEY")),
		SuccessHandler: func(c echo.Context) {
			claims := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)
			userId := uint(claims["id"].(float64))

			users := query.User
			user, err := users.Where(users.ID.Eq(userId)).
				Preload(users.Wallet).
				First()

			if err == nil {
				c.Set("user", user)
			}
		},
	}
}
