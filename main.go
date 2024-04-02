package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	config "pay-bills/app"
	"pay-bills/app/helpers"
	"pay-bills/database"
	"pay-bills/routes"
)

func init() {
	config.Setup()

	database.Connect()

	helper.CustomRules()
}

func main() {
	e := echo.New()

	routes.Api(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", viper.GetString("APP_HOST"), viper.GetInt("APP_PORT"))))
}
