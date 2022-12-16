package api

import (
	"fmt"
	"log"
	"marcelofelixsalgado/financial-web/api/controllers/credentials"
	"marcelofelixsalgado/financial-web/api/controllers/health"
	"marcelofelixsalgado/financial-web/api/controllers/user"
	"marcelofelixsalgado/financial-web/api/routes"
	"marcelofelixsalgado/financial-web/configs"
	"marcelofelixsalgado/financial-web/utils"
	"net/http"

	"github.com/gorilla/mux"

	userCreate "marcelofelixsalgado/financial-web/pkg/usecase/user/create"

	userCredentialsCreate "marcelofelixsalgado/financial-web/pkg/usecase/credentials/create"
	userCredentialsLogin "marcelofelixsalgado/financial-web/pkg/usecase/credentials/login"
)

func NewServer() *mux.Router {
	// Load environment variables
	configs.Load()

	// Load HTML templates
	utils.LoadTemplates()

	userCredentialsRoutes := setupUserCredentialsRoutes()
	userRoutes := setupUserRoutes()
	healthRoutes := setupHealthRoutes()

	// Setup all routes
	routes := routes.NewRoutes(userCredentialsRoutes, userRoutes, healthRoutes)

	router := routes.SetupRoutes()
	return router
}

func Run(router *mux.Router) {
	port := fmt.Sprintf(":%d", configs.HttpPort)

	log.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func setupUserRoutes() user.UserRoutes {

	// setup Use Cases (services)
	createUseCase := userCreate.NewCreateUseCase()

	// setup router handlers
	userHandler := user.NewUserHandler(createUseCase)

	// setup routes
	userRoutes := user.NewUserRoutes(userHandler)

	return userRoutes
}

func setupUserCredentialsRoutes() credentials.UserCredentialsRoutes {

	// setup Use Cases (services)
	createUseCase := userCredentialsCreate.NewCreateUseCase()
	loginUseCase := userCredentialsLogin.NewLoginUseCase()

	// setup router handlers
	userCredentialsHandler := credentials.NewUserCredentialsHandler(createUseCase, loginUseCase)

	// setup routes
	userCredentialsRoutes := credentials.NewUserCredentialsRoutes(userCredentialsHandler)

	return userCredentialsRoutes
}

func setupHealthRoutes() health.HealthRoutes {
	// setup router handlers
	healthHandler := health.NewHealthHandler()

	// setup routes
	healthRoutes := health.NewHealthRoutes(healthHandler)

	return healthRoutes
}
