package provider

import (
	"fmt"
	"log"

	"github.com/qlixes/helpdesk/internal/config"
	"github.com/spf13/viper"
)

type Config struct {
	App *config.App
	Db  *config.Db
}

type ConfigMethod interface {
	GetAppAddress() string
	GetAppName() string
	GetPgsqlConnection() string
	GetMysqlConnection() string
	GetAppTimezone() string
}

var Cfg = newConfig()

func newConfig() *Config {

	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Config not found : %s \n", err.Error())
	}

	return &Config{
		App: &config.App{
			Host:     viper.GetString("APP_HOST"),
			Port:     viper.GetInt("APP_PORT"),
			Name:     viper.GetString("APP_NAME"),
			TimeZone: viper.GetString("APP_TIMEZONE"),
		},
		Db: &config.Db{
			Host: viper.GetString("DB_HOST"),
			Port: viper.GetInt("DB_PORT"),
			User: viper.GetString("DB_USER"),
			Pass: viper.GetString("DB_PASS"),
			Name: viper.GetString("DB_NAME"),
		},
	}
}

func (c *Config) GetAppAddress() string {
	return fmt.Sprintf("%s:%d", c.App.Host, c.App.Port)
}

func (c *Config) GetPgsqlConnection() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable&TimeZone=Asia/Jakarta",
		c.Db.User,
		c.Db.Pass,
		c.Db.Host,
		c.Db.Port,
		c.Db.Name,
	)
}

func (c *Config) GetMysqlConnection() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Db.User,
		c.Db.Pass,
		c.Db.Host,
		c.Db.Port,
		c.Db.Name,
	)
}

func (c *Config) GetAppName() string {
	return c.App.Name
}

func (c *Config) GetAppTimezone() string {
	return c.App.TimeZone
}
