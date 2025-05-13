package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/csrf"
	"github.com/qlixes/helpdesk/internal/provider"
	"github.com/qlixes/helpdesk/internal/router"
)

func main() {

	app := fiber.New(fiber.Config{
		CaseSensitive:     true,
		StrictRouting:     true,
		AppName:           provider.Cfg.GetAppName(),
		Immutable:         true,
		ReduceMemoryUsage: true,
	})

	app.Use(compress.New(compress.Config{}))
	app.Use(cors.New())
	app.Use(csrf.New())

	router.WebRouter(app)
	router.ApiRouter(app)

	err := app.Listen(provider.Cfg.GetAppAddress(), fiber.ListenConfig{
		EnablePrefork:         true,
		DisableStartupMessage: false,
	})

	if err != nil {
		log.Fatalf("Unable run service : %s \n", err.Error())
	}
}
