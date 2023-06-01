package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/lctzz540/chessaiserver/middlewares"
	"github.com/lctzz540/chessaiserver/routes"
	"github.com/lctzz540/chessaiserver/utils"
)

func main() {
	utils.SetupRedis()
	app := fiber.New()

	middlewares.CorsMiddleware(app)

	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return strings.Contains(c.Route().Path, "/ws")
		},
	}))
	routes.Websocket(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello")
	})

	app.Listen(":3000")
}
