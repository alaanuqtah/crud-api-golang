package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	err := godotenv.Load() //load the .env file

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
