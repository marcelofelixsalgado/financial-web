package settings

import (
	"log"

	"github.com/codingconcepts/env"
	"github.com/joho/godotenv"
)

// ConfigType struct to resolve env vars
type ConfigType struct {
	// Application name
	AppName string `env:"APP_NAME" default:"financial-web"`

	// Current environment
	Environment string `env:"ENVIRONMENT" default:"development"`

	// HTTP Port to expose the Web API
	WebHttpPort int `env:"WEB_HTTP_PORT" default:"8000"`

	// User API address and port
	UserApiURL string `env:"USER_API_URL"`

	// Period API address and port
	PeriodApiURL string `env:"PERIOD_API_URL"`

	// Used to authenticate the cookie
	HashKey []byte `env:"HASH_KEY"`

	// Used to encrypt cookie data
	BlockKey []byte `env:"BLOCK_KEY"`

	// Time waiting to server shutdown
	ServerCloseWait int `env:"SERVER_CLOSEWAIT" default:"10"`
}

var Config ConfigType

// InitConfigs initializes the environment settings
func Load() {
	// load .env (if exists)
	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found")
	}

	// bind env vars
	if err := env.Set(&Config); err != nil {
		log.Fatal(err)
	}
}
