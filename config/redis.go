package config

import (
	"os"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func ConnectRedisClient() {
	client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
		Protocol: 2,
	})
}

func GetRedisClient() *redis.Client {
	return client
}
