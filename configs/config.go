package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// HTTP Port to expose the Web API
	WebHttpPort = 0

	// User API address and port
	UserApiURL = ""

	// Period API address and port
	PeriodApiURL = ""

	// Used to authenticate the cookie
	HashKey []byte

	// Used to encrypt cookie data
	BlockKey []byte
)

// Load global parameters from environment variables
func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatalf("Error trying to load the environment variables: %v", err)
	}

	WebHttpPort, err = strconv.Atoi(os.Getenv("WEB_HTTP_PORT"))
	if err != nil {
		log.Fatalf("Could not find the WEB_HTTP_PORT environment variable: %v", err)
	}

	UserApiURL = os.Getenv("USER_API_URL")
	if UserApiURL == "" {
		log.Fatalf("Could not find the USER_API_URL environment variable: %v", err)
	}

	PeriodApiURL = os.Getenv("PERIOD_API_URL")
	if UserApiURL == "" {
		log.Fatalf("Could not find the PERIOD_API_URL environment variable: %v", err)
	}

	HashKey = []byte(os.Getenv("HASH_KEY"))
	if len(HashKey) == 0 {
		log.Fatalf("Could not find the HASH_KEY environment variable")
	}

	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
	if len(BlockKey) == 0 {
		log.Fatalf("Could not find the BLOCK_KEY environment variable")
	}
}
