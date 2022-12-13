package routes

import (
	"marcelofelixsalgado/financial-web/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:                    "/register",
		Method:                 http.MethodGet,
		Function:               controllers.LoadUserRegisterPage,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Function:               controllers.CreateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/register/credentials",
		Method:                 http.MethodPost,
		Function:               controllers.LoadUserRegisterCredentialsPage,
		RequiresAuthentication: false,
	},
}
