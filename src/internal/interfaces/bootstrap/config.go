package bootstrap

import "os"

type Config struct {
	AppPort string
	AppHost string
	AppName string

	DbHost string
	DbPort string
	DbUser string
	DbPass string
	DbName string

	MailHost string
	MailPort string
	MailUser string
	MailPass string
	MailName string
	MailFrom string
}

func NewConfig() *Config {

	return &Config{
		AppPort: os.Getenv("APP_PORT"),
		AppHost: os.Getenv("APP_HOST"),
		AppName: os.Getenv("APP_NAME"),

		DbHost: os.Getenv("DB_HOST"),
		DbPort: os.Getenv("DB_PORT"),
		DbUser: os.Getenv("DB_USER"),
		DbPass: os.Getenv("DB_PASS"),
		DbName: os.Getenv("DB_NAME"),

		MailHost: os.Getenv("MAIL_HOST"),
		MailPort: os.Getenv("MAIL_PORT"),
		MailUser: os.Getenv("MAIL_USER"),
		MailPass: os.Getenv("MAIL_PASS"),
		MailFrom: os.Getenv("MAIL_FROM"),
		MailName: os.Getenv("MAIL_NAME"),
	}
}
