package utils

import (
	"os"

	"github.com/redis/go-redis/v9"
)

func Cache() *redis.Client {
    redisAddress := os.Getenv("REDIS_ADDRESS")
    redisPassword := os.Getenv("REDIS_PASSWORD")
	client := redis.NewClient(&redis.Options{
        Addr:	  redisAddress,
        Password: redisPassword,
        DB:		  0,
    })

	return client
}