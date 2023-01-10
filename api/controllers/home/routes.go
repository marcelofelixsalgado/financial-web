package home

import (
	"marcelofelixsalgado/financial-web/api/controllers"
	"net/http"
)

var homeBasepath = "/home"

type HomeRoutes struct {
	homeHandler IHomeHandler
}

func NewHomeRoutes(homeHandler IHomeHandler) HomeRoutes {
	return HomeRoutes{
		homeHandler: homeHandler,
	}
}

func (homeRoutes *HomeRoutes) HomeRouteMapping() (string, []controllers.Route) {

	return homeBasepath, []controllers.Route{
		{
			URI:                    "",
			Method:                 http.MethodGet,
			Function:               homeRoutes.homeHandler.Home,
			RequiresAuthentication: true,
		},
	}
}
