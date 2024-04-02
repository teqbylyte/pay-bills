package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

// Setup - Set up the configuration/environment variables and timezone
func Setup() {
	configureTimezone()

	viper.SetConfigType("toml")
	viper.SetConfigName(".app")

	//viper.AddConfigPath(fmt.Sprintf("%s/prod/pay-bills", homeDir))
	viper.AddConfigPath("./")

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
}

// Configure timezone for application
func configureTimezone() {
	timezone := "Africa/Lagos"

	_ = os.Setenv("TZ", timezone)

	location, err := time.LoadLocation(timezone)

	if err != nil {
		fmt.Printf("Err loading location: %v\n", err)
	}

	time.Local = location
}
