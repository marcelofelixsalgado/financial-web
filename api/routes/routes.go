package routes

import (
	"marcelofelixsalgado/financial-web/api/controllers"
	"marcelofelixsalgado/financial-web/api/controllers/credentials"
	"marcelofelixsalgado/financial-web/api/controllers/health"
	"marcelofelixsalgado/financial-web/api/controllers/home"
	"marcelofelixsalgado/financial-web/api/controllers/period"
	"marcelofelixsalgado/financial-web/api/controllers/user"
	"marcelofelixsalgado/financial-web/api/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// func SetupRoutes() *mux.Router {
// 	router := mux.NewRouter()

// 	// user routes
// 	for _, route := range register.UserRoutes {
// 		router.HandleFunc(route.URI,
// 			middlewares.Logger(
// 				route.Function)).Methods(route.Method)
// 	}

// 	// login routes
// 	for _, route := range login.LoginRoutes {
// 		router.HandleFunc(route.URI,
// 			middlewares.Logger(
// 				route.Function)).Methods(route.Method)
// 	}

// 	// Setting the path for static files in assets folder
// 	fileServer := http.FileServer(http.Dir("./assets/"))
// 	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

// 	return router
// }

type Routes struct {
	userCredentialRoutes credentials.UserCredentialsRoutes
	userRoutes           user.UserRoutes
	homeRoutes           home.HomeRoutes
	periodRoutes         period.PeriodRoutes
	healthRoutes         health.HealthRoutes
}

func NewRoutes(userCredentialRoutes credentials.UserCredentialsRoutes, userRoutes user.UserRoutes, homeRoutes home.HomeRoutes, periodRoutes period.PeriodRoutes, healthRoutes health.HealthRoutes) *Routes {
	return &Routes{
		userCredentialRoutes: userCredentialRoutes,
		userRoutes:           userRoutes,
		homeRoutes:           homeRoutes,
		periodRoutes:         periodRoutes,
		healthRoutes:         healthRoutes,
	}
}

func (routes *Routes) SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	// router.Use(middlewares.ResponseFormatMiddleware)

	// user credentials routes
	setupRoute(router, routes.userCredentialRoutes.UserCredentialsRouteMapping())

	// user routes
	setupRoute(router, routes.userRoutes.UserRouteMapping())

	// home routes
	setupRoute(router, routes.homeRoutes.HomeRouteMapping())

	// period routes
	setupRoute(router, routes.periodRoutes.PeriodRouteMapping())

	// health routes
	setupRoute(router, routes.healthRoutes.HealthRouteMapping())

	// Setting the path for static files in assets folder
	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}

func setupRoute(router *mux.Router, routes []controllers.Route) {
	for _, route := range routes {
		if route.RequiresAuthentication {
			router.HandleFunc(route.URI,
				middlewares.Logger(
					middlewares.Authenticate(route.Function))).Methods(route.Method)
		} else {
			router.HandleFunc(route.URI,
				middlewares.Logger(route.Function)).Methods(route.Method)

		}
	}
}
