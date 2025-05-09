package router

import "github.com/gofiber/fiber/v3"

func WebRouter(r *fiber.App) {
	r.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello world")
	})
}
