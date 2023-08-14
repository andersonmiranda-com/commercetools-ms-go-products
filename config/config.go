package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Getenv(key string) string {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Fatalln("Error loading .env")
	}
	return os.Getenv(key)

}
