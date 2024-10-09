package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config holds the configuration values
type Config struct {
	DataBaseUrl string
	HostPort    int
	SecretKey   string
}

// LoadConfig loads environment variables using Viper
func LoadConfig() Config {
	// Set the file name of the configuration file without the extension
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	// Set the path to look for the configuration file
	viper.AddConfigPath(".") // or specify a specific directory, e.g., "/path/to/env/"

	// Enable viper to read environment variables
	viper.AutomaticEnv()

	// Read in the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Return the configuration struct with values from Viper
	return Config{
		DataBaseUrl: viper.GetString("DATABASE_URL"),
		SecretKey:   viper.GetString("SECRET_KEY"),
		HostPort:    viper.GetInt("APPLICATION_PORT"),
	}
}
