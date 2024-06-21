package main

import (
	"fmt"
	"os"

	"aalstek.com/m/v2/config"
	"aalstek.com/m/v2/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("./", "/static")
	app.Static("./", "/routes")

	configuration, err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	router.Route(app)

	app.Listen(config.FormatPort(configuration.Server.Port))
}
