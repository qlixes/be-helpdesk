package bootstrap

import (
	"os"

	"github.com/qlixes/be-helpdesk/internal/config"
)

type Config struct {
	app   *config.App
	db    *config.Db
	cache *config.Cache
	queue *config.Queue
	mail  *Config.Mail
}

type ConfigManager interface {
	GetAppPort() string
}

func NewConfig() *Config {

	return &Config{
		app: &config.App{
			Port: os.Getenv("APP_PORT"),
			Host: os.Getenv("APP_HOST"),
			Name: os.Getenv("APP_NAME"),
		},

		db: &config.Db{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASS"),
			Name:     os.Getenv("DB_NAME"),
		},

		mail: &config.Mail{
			Host:     os.Getenv("MAIL_HOST"),
			Port:     os.Getenv("MAIL_PORT"),
			User:     os.Getenv("MAIL_USER"),
			Password: os.Getenv("MAIL_PASS"),
			From:     os.Getenv("MAIL_FROM"),
			Name:     os.Getenv("MAIL_NAME"),
		},
	}
}

func (c *Config) GetAppPort() string {
	return c.app.Port
}
