package models

import (
	redis "gopkg.in/redis.v3"
)

var rdx *redis.Client

func GetRedisClient() *redis.Client {
	if rdx != nil {
		return rdx
	}

	rdx = redis.NewClient(&redis.Options{
		Addr:     REDIS_ADDR,
		Password: REDIS_PASSWORD,
		DB:       0,
	})
	return rdx
}
