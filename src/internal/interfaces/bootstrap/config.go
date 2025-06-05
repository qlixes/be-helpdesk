package bootstrap

import "github.com/spf13/viper"

type Config struct {
	AppPort int
	AppHost string
	AppName string
}

func NewConfig() *Config {
	return &Config{
		AppPort: viper.GetInt("APP_PORT"),
		AppHost: viper.GetString("APP_HOST"),
		AppName: viper.GetString("APP_NAME"),
	}
}
