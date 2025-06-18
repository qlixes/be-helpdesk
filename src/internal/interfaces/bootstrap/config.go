package bootstrap

import (
	"fmt"
	"os"

	"github.com/qlixes/be-helpdesk/internal/config"
)

type Config struct {
	app      *config.App
	db       *config.Db
	cache    *config.Cache
	queue    *config.Queue
	mail     *config.Mail
	provider *config.Provider
}

type ConfigManager interface {
	GetAppHost() string
	GetAppPort() string
	GetAppName() string
	GetAppKey() string

	GetDbEngine() (string, string)
	GetDbHost() string
	GetDbPort() string
	GetDbUser() string
	GetDbPassword() string
	GetDbName() string
}

func NewConfig() *Config {

	return &Config{
		app: &config.App{
			Host: os.Getenv("APP_HOST"),
			Port: os.Getenv("APP_PORT"),
			Name: os.Getenv("APP_NAME"),
			Key:  os.Getenv("APP_KEY"),
		},

		db: &config.Db{
			Engine:   os.Getenv("DB_ENGINE"),
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

		provider: &config.Provider{},
	}
}

func (c *Config) GetAppHost() string {
	return c.app.Host
}

func (c *Config) GetAppPort() string {
	return c.app.Port
}

func (c *Config) GetAppHost() string {
	return c.app.Host
}

func (c *Config) GetAppName() string {
	return c.app.Name
}

func (c *Config) GetAppKey() string {
	return c.app.Key
}

func (c *Config) GetDbEngine() (string, string) {

	if c.db.Engine == "" {
		panic("could not detect db engine")
	}

	dsn := ""

	switch c.db.Engine {
	case "pgsql":
		dsn = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
			c.db.User,
			c.db.Password,
			c.db.Host,
			c.db.Port,
			c.db.Name,
		)
	default:
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			c.db.User,
			c.db.Password,
			c.db.Host,
			c.db.Port,
			c.db.Name,
		)
	}

	return c.db.Engine, dsn
}

func (c *Config) GetDbHost() string {
	return c.db.Host
}

func (c *Config) GetDbPort() string {
	return c.db.Port
}

func (c *Config) GetDbUser() string {
	return c.db.User
}

func (c *Config) GetDbPassword() string {
	return c.db.Password
}

func (c *Config) GetDbName() string {
	return c.db.Name
}
