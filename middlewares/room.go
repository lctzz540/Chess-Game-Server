package middlewares

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/lctzz540/chessaiserver/models"
)

func WebsocketConnection() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	}
}
func ConnectToRoomID(connections map[string]*models.ConnectionCount) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		connCount, ok := connections[id]
		if ok && connCount.Count >= connCount.MaxConn {
			log.Printf("Rejected WebSocket connection: %s\n", id)
			return fiber.ErrTooManyRequests
		}

		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}

		return fiber.ErrUpgradeRequired
	}
}
