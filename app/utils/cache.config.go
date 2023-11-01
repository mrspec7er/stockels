package utils

import (
	"os"

	"github.com/redis/go-redis/v9"
)

func Cache() *redis.Client {
    redisAddress := os.Getenv("REDIS_ADDRESS")
    redisPassword := os.Getenv("REDIS_PASSWORD")
    redisUsername := os.Getenv("REDIS_USERNAME")
	client := redis.NewClient(&redis.Options{
        Addr:	  redisAddress,
        Password: redisPassword,
        Username: redisUsername,
        DB:		  0,
     
    })

	return client
}