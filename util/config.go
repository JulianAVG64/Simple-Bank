package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configurations of the application.
// The values are read by viper from a config file or enviroment variables
type Config struct {
	DBDriver                   string        `mapstructure:"DB_DIVER"`
	DBSource                   string        `mapstructure:"DB_SOURCE"`
	ServerAddress              string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey          string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration        time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshAccessTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

// LoadConfig reads configuration from file or enviroment variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
