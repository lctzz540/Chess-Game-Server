package routes

import (
	"context"
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/lctzz540/chessaiserver/handlers"
	"github.com/lctzz540/chessaiserver/middlewares"
	"github.com/lctzz540/chessaiserver/models"
)

var connections = make(map[string]*models.ConnectionCount)

func Websocket(app *fiber.App, redisClient *redis.Client) {
	app.Use("/ws", middlewares.WebSocketConnection())

	app.Get("/ws/:id", middlewares.ConnectToRoomID(connections), websocket.New(func(c *websocket.Conn) {
		roomId := c.Params("id")
		ctx := context.Background()
		chessGame := handlers.StartGame(ctx, roomId, redisClient)
		if chessGame == nil {
			return
		}

		connCount, ok := connections[roomId]
		if !ok {
			connCount = &models.ConnectionCount{Count: 0, MaxConn: 2}
			connections[roomId] = connCount
		}
		if connCount.Count >= connCount.MaxConn {
			log.Printf("Rejected WebSocket connection: %s\n", roomId)
			c.Close()
			return
		}
		connCount.Count++

		log.Println("New WebSocket connection:", roomId)

		defer func() {
			connCount.Count--
			log.Println("WebSocket connection closed:", roomId)
		}()

		for {
			_, p, err := c.ReadMessage()
			if err != nil {
				log.Println("WebSocket read error:", err)
				break
			}
			var move models.Move
			if err := json.Unmarshal(p, &move); err != nil {
				log.Println("Failed to unmarshal move:", err)
				continue
			}
			err = handlers.MakeMove(ctx, chessGame, roomId, &move, redisClient)
			if err != nil {
				log.Println("Failed to make move:", err)
				return
			}

			log.Printf("Received message from client %s: %s\n", roomId, string(p))
		}
	}))
}
