package landing

import "github.com/gofiber/fiber/v2"

func Get(context *fiber.Ctx) error {
	return context.Render("routes/landing/landing", fiber.Map{
		"value": "World",
	})
}
