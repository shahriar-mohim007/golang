package config

import (
	"github.com/spf13/viper"
)

// Config holds the configuration values
type Config struct {
	DataBaseUrl    string
	HostPort       int
	SecretKey      string
	Rps            float64
	Burst          int
	LimiterEnabled bool
}

// LoadConfig loads environment variables using Viper
func LoadConfig() Config {

	viper.AutomaticEnv()

	return Config{
		DataBaseUrl:    viper.GetString("DATABASE_URL"),
		SecretKey:      viper.GetString("SECRET_KEY"),
		HostPort:       viper.GetInt("APPLICATION_PORT"),
		Rps:            viper.GetFloat64("limiter_rps"),
		Burst:          viper.GetInt("limiter_burst"),
		LimiterEnabled: viper.GetBool("limiter_enabled"),
	}
}
