package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvConfig(key string) string {

	err := godotenv.Load("app.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
