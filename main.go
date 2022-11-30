package main

import (
	"marcelofelixsalgado/financial-web/configs"
	"marcelofelixsalgado/financial-web/utils"
)

func main() {
	// Load environment variables
	configs.Load()

	// Load HTML templates
	utils.LoadTemplates()

	// Start HTTP Server
	startServer()
}
