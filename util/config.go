package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configurations of the application.
// The values are read by viper from a config file or enviroment variables
type Config struct {
	Enviroment                 string        `mapstructure:"ENVIROMENT"`
	DBSource                   string        `mapstructure:"DB_SOURCE"`
	MigrationURL               string        `mapstructure:"MIGRATION_URL"`
	RedisAddress               string        `mapstructure:"REDIS_ADDRESS"`
	HTTPServerAddress          string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress          string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey          string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration        time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshAccessTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	EmailSenderName            string        `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress         string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword        string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
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
