package redisdb

import (
	"os"
	"log"

	"github.com/go-redis/redis/v8" 
	"golang.org/x/net/context"
)

var Client *redis.Client

func ConnectRedis() *redis.Client {
	// Correct the environment variable to use REDIS_ADDR instead of REDIS_PORT
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	// Create a new Redis client
	Client = redis.NewClient(&redis.Options{
		Addr:     redisAddr,  // The full address of the Redis server (e.g., "localhost:6379")
		Password: redisPassword, // Redis password if set, otherwise it's empty
		DB:       0, // Default DB to use
	})

	// Ping the Redis server to check the connection
	ctx := context.Background()
	_, err := Client.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis: ", err)
	}

	return Client
}
