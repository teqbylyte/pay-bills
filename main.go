package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	config "martpay/app"
	db "martpay/database"
	"martpay/routes"
)

func init() {
	config.Setup()

	db.Connect()
}

func main() {
	e := echo.New()

	routes.Api(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", viper.GetString("APP_HOST"), viper.GetInt("APP_PORT"))))
}
