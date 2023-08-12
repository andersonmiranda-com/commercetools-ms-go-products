package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Getenv(key string) string {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	return os.Getenv(key)

}
