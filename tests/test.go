package tests

import (
	"log"
	"os"
)

func init() {
	accessToken := os.Getenv("HS_ACCESS_TOKEN")
	if accessToken == "" {
		log.Fatal("HS_ACCESS_TOKEN environment variable is not set")
	}
}
