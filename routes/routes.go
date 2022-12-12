package routes

import (
	"marcelofelixsalgado/financial-web/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI                    string
	Method                 string
	Function               func(w http.ResponseWriter, r *http.Request)
	RequiresAuthentication bool
}

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// login routes
	for _, route := range loginRoutes {
		router.HandleFunc(route.URI,
			middlewares.Logger(
				route.Function)).Methods(route.Method)
	}

	// Setting the path for static files in assets folder
	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
