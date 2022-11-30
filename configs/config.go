package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// HTTP Port to expose the API
	HttpPort = 0
)

// Load global parameters from environment variables
func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatalf("Error trying to load the environment variables: %v", err)
	}

	HttpPort, err = strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		log.Fatalf("Could not find the HTTP_PORT environment variable: %v", err)
	}
}
