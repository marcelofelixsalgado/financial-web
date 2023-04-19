package main

import "github.com/marcelofelixsalgado/financial-web/api"

func main() {
	// Start HTTP Server
	server := api.NewServer()
	server.Run()
}
