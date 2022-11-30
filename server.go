package main

import (
	"fmt"
	"log"
	"marcelofelixsalgado/financial-web/configs"
	"marcelofelixsalgado/financial-web/routes"

	"net/http"
)

func startServer() {
	port := fmt.Sprintf(":%d", configs.HttpPort)

	router := routes.SetupRoutes()

	fmt.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
