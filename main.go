package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"log"
	"martpay/routes"
	"os"
	"time"
)

func init() {
	_ = os.Setenv("TZ", "Africa/Lagos")
	location, err := time.LoadLocation("Africa/Lagos")
	if err != nil {
		fmt.Printf("Err loading location: %v\n", err)
	}
	time.Local = location

	//homeDir, _ := os.UserHomeDir()
	viper.SetConfigType("toml")
	viper.SetConfigName(".martpay")

	//viper.AddConfigPath(fmt.Sprintf("%s/prod/martpay", homeDir))
	viper.AddConfigPath("./")

	// Find and read the config file
	err = viper.ReadInConfig()
	fmt.Println(err)

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	// Init Database
	//database.Database{}.Init()
}

func main() {
	e := echo.New()

	routes.Api(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", viper.GetString("APP_HOST"), viper.GetInt("APP_PORT"))))
}
