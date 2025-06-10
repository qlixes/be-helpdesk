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
	mail  *config.Mail
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

		queue: &config.Queue{
			Host:     os.Getenv("QUEUE_HOST"),
			Port:     os.Getenv("QUEUE_PORT"),
			User:     os.Getenv("QUEUE_USER"),
			Password: os.Getenv("QUEUE_PASS"),
			Vhost:    os.Getenv("QUEUE_VHOST"),
		},

		cache: &config.Cache{
			Host:     os.Getenv("CACHE_HOST"),
			Port:     os.Getenv("CACHE_PORT"),
			User:     os.Getenv("CACHE_USER"),
			Password: os.Getenv("CACHE_PASS"),
			Index:    os.Getenv("CACHE_INDEX"),
		},
	}
}

func (c *Config) GetAppPort() string {
	return c.app.Port
}
