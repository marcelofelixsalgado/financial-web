package api

import (
	"fmt"
	"log"
	"marcelofelixsalgado/financial-web/api/controllers/credentials"
	"marcelofelixsalgado/financial-web/api/controllers/health"
	"marcelofelixsalgado/financial-web/api/controllers/home"
	"marcelofelixsalgado/financial-web/api/controllers/user"
	"marcelofelixsalgado/financial-web/api/cookies"
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

	// Configure cookies
	cookies.Configure()

	userCredentialsRoutes := setupUserCredentialsRoutes()
	userRoutes := setupUserRoutes()
	homeRoutes := setupHomeRoutes()
	healthRoutes := setupHealthRoutes()

	// Setup all routes
	routes := routes.NewRoutes(userCredentialsRoutes, userRoutes, homeRoutes, healthRoutes)

	router := routes.SetupRoutes()
	return router
}

func Run(router *mux.Router) {
	port := fmt.Sprintf(":%d", configs.WebHttpPort)

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

func setupHomeRoutes() home.HomeRoutes {

	// setup router handlers
	homeHandler := home.NewHomeHandler()

	// setup routes
	homeRoutes := home.NewHomeRoutes(homeHandler)

	return homeRoutes
}

func setupHealthRoutes() health.HealthRoutes {
	// setup router handlers
	healthHandler := health.NewHealthHandler()

	// setup routes
	healthRoutes := health.NewHealthRoutes(healthHandler)

	return healthRoutes
}
