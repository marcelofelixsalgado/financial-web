package home

import (
	"marcelofelixsalgado/financial-web/api/controllers"
	"net/http"
)

type HomeRoutes struct {
	homeHandler IHomeHandler
}

func NewHomeRoutes(homeHandler IHomeHandler) HomeRoutes {
	return HomeRoutes{
		homeHandler: homeHandler,
	}
}

func (homeRoutes *HomeRoutes) HomeRouteMapping() []controllers.Route {

	return []controllers.Route{
		{
			URI:                    "/home",
			Method:                 http.MethodGet,
			Function:               controllers.LoadUserHomePage,
			RequiresAuthentication: true,
		},
	}
}
