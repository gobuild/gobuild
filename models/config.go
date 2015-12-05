package models

import "os"

var (
	GITHUB_TOKEN   = os.Getenv("GITHUB_TOKEN")
	REDIS_ADDR     = os.Getenv("REDIS_ADDR")
	REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
)

func init() {
	if REDIS_ADDR == "" {
		REDIS_ADDR = "localhost:6379"
	}
}
