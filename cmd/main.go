package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/qlixes/helpdesk/internal/provider"
	"github.com/qlixes/helpdesk/internal/router"
)

func main() {

	app := fiber.New(fiber.Config{})

	router.WebRouter(app)
	router.ApiRouter(app)

	err := app.Listen(provider.Cfg.GetAppAddress())

	if err != nil {
		log.Fatalf("Unable run service : %s \n", err.Error())
	}
}
