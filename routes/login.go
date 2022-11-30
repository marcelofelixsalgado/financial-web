package routes

import (
	"marcelofelixsalgado/financial-web/controllers"
	"net/http"
)

var loginRoutes = []Route{
	{
		URI:                    "/",
		Method:                 http.MethodGet,
		Function:               controllers.LoadLoginScreen,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/login",
		Method:                 http.MethodGet,
		Function:               controllers.LoadLoginScreen,
		RequiresAuthentication: false,
	},
}
