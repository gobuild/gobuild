package models

import "os"

var (
	GITHUB_TOKEN   = os.Getenv("GITHUB_TOKEN")
	REDIS_ADDR     = os.Getenv("REDIS_ADDR")
	REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
	MYSQL_URI      = os.Getenv("MYSQL_URI")
	CDN_URL_BASE   = os.Getenv("CDN_URL_BASE")
)

func init() {
	if REDIS_ADDR == "" {
		REDIS_ADDR = "localhost:6379"
	}
	if CDN_URL_BASE == "" {
		CDN_URL_BASE = "http://dn-gobuild5.qbox.me/gorelease/"
	}
}
