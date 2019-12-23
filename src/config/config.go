package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// PORT set begin param
var PORT = 0

// Load config file
func Load() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		PORT = 9000
	}
}
