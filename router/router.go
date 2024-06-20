package router

import (
	"aalstek.com/m/v2/routes/landing"
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) error {
	app.Get("/", landing.Get)

	return nil
}
