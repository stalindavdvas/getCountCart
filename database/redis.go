package database

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// InitRedis initializes and returns a Redis client
func InitRedis() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     "52.5.28.74:6379", // Dirección de Redis
		Password: "",                // Contraseña (si aplica)
		DB:       0,
	})

	// Test the connection
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	fmt.Println("Successfully connected to Redis")
	return client
}
