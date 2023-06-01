package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

var RedisClient *redis.Client

func SetupRedis() {
	// Create a new Redis client
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // Redis container hostname and port
		Password: "",               // Redis server password, if any
		DB:       0,                // Redis database index
	})

	// Ping the Redis server to test the connection
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	fmt.Println("Connected to Redis")
}

func CloseRedis() {
	// Close the Redis connection when done
	if err := RedisClient.Close(); err != nil {
		log.Println("Error closing Redis connection:", err)
	}
}

func RedisMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("redis", RedisClient)
		return c.Next()
	}
}
