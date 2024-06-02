package helpers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Load the environment variables
	// Find .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	log.Println("Loaded .env file")
}
