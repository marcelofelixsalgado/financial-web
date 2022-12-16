package main

import "marcelofelixsalgado/financial-web/api"

func main() {
	// Start HTTP Server
	router := api.NewServer()
	api.Run(router)
}
