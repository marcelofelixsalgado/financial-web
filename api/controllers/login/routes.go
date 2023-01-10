package login

import (
	"marcelofelixsalgado/financial-web/api/controllers"
	"net/http"
)

var credentialsBasePath = ""

type LoginRoutes struct {
	LoginHandler ILoginHandler
}

func NewLoginRoutes(LoginHandler ILoginHandler) LoginRoutes {
	return LoginRoutes{
		LoginHandler: LoginHandler,
	}
}

func (LoginRoutes *LoginRoutes) LoginRouteMapping() (string, []controllers.Route) {

	return credentialsBasePath, []controllers.Route{
		{
			URI:                    "/",
			Method:                 http.MethodGet,
			Function:               LoginRoutes.LoginHandler.LoadLoginPage,
			RequiresAuthentication: false,
		},
		{
			URI:                    "/login",
			Method:                 http.MethodGet,
			Function:               LoginRoutes.LoginHandler.LoadLoginPage,
			RequiresAuthentication: false,
		},
		{
			URI:                    "/login",
			Method:                 http.MethodPost,
			Function:               LoginRoutes.LoginHandler.Login,
			RequiresAuthentication: false,
		},
	}
}
