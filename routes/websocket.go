package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type ConnectionCount struct {
	Count   int
	MaxConn int
}

var connections = make(map[string]*ConnectionCount)

func Websocket(app *fiber.App) {
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Use("/ws/:id", func(c *fiber.Ctx) error {
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
	})

	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		id := c.Params("id")

		connCount, ok := connections[id]
		if !ok {
			connCount = &ConnectionCount{Count: 0, MaxConn: 2}
			connections[id] = connCount
		}
		if connCount.Count >= connCount.MaxConn {
			log.Printf("Rejected WebSocket connection: %s\n", id)
			c.Close()
			return
		}
		connCount.Count++

		log.Println("New WebSocket connection:", id)

		defer func() {
			connCount.Count--
			log.Println("WebSocket connection closed:", id)
		}()

		for {
			_, p, err := c.ReadMessage()
			if err != nil {
				log.Println("WebSocket read error:", err)
				break
			}

			log.Printf("Received message from client %s: %s\n", id, string(p))
		}
	}))
}
