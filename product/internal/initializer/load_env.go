package initializer

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("./internal/config/.env")

	if err != nil {
		log.Fatalln("failed to load .env file")
	}
}
