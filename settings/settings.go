package settings

import (
	"log"
	"os"

	"github.com/codingconcepts/env"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
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

	// Transaction Type, Category and SubCategory APIs address and port
	CategoryApiURL string `env:"CATEGORY_API_URL"`

	// Period API address and port
	PeriodApiURL string `env:"PERIOD_API_URL"`

	// Balance API address and port
	BalanceApiURL string `env:"BALANCE_API_URL"`

	// Used to authenticate the cookie
	HashKey []byte `env:"HASH_KEY"`

	// Used to encrypt cookie data
	BlockKey []byte `env:"BLOCK_KEY"`

	// Time waiting to server shutdown
	ServerCloseWait int `env:"SERVER_CLOSEWAIT" default:"10"`

	// Log files
	LogAccessFile string `env:"LOG_ACCESS_FILE" default:"./access.log"`
	LogAppFile    string `env:"LOG_APP_FILE" default:"./app.log"`
	LogLevel      string `env:"LOG_LEVEL" default:"INFO"`
}

var Config ConfigType

// InitConfigs initializes the environment settings
func Load() {
	// load .env (if exists)
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Println("No .env file found")
	// }
	loadDotEnv()

	// bind env vars
	if err := env.Set(&Config); err != nil {
		log.Fatal(err)
	}

	if _, err := logrus.ParseLevel(Config.LogLevel); err != nil {
		log.Fatal(err)
	}
}

func loadDotEnv() {
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			logrus.Fatal("Error loading .env file")
		}
		logrus.Println("Using .env file")
	}
}
