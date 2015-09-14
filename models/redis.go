package models

import (
	"os"

	"gopkg.in/redis.v3"
)

var rdx *redis.Client

func GetRedisClient() *redis.Client {
	if rdx != nil {
		return rdx
	}
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}
	rdx = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	return rdx
}
