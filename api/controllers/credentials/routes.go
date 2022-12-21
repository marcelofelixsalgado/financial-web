package credentials

import (
	"marcelofelixsalgado/financial-web/api/controllers"
	"net/http"
)

type UserCredentialsRoutes struct {
	userCredentialsHandler IUserCredentialsHandler
}

func NewUserCredentialsRoutes(userCredentialsHandler IUserCredentialsHandler) UserCredentialsRoutes {
	return UserCredentialsRoutes{
		userCredentialsHandler: userCredentialsHandler,
	}
}

func (userCredentialsRoutes *UserCredentialsRoutes) UserCredentialsRouteMapping() []controllers.Route {

	return []controllers.Route{
		{
			URI:                    "/",
			Method:                 http.MethodGet,
			Function:               userCredentialsRoutes.userCredentialsHandler.LoadLoginPage,
			RequiresAuthentication: false,
		},
		{
			URI:                    "/login",
			Method:                 http.MethodGet,
			Function:               userCredentialsRoutes.userCredentialsHandler.LoadLoginPage,
			RequiresAuthentication: false,
		},
		{
			URI:                    "/register/credentials",
			Method:                 http.MethodGet,
			Function:               userCredentialsRoutes.userCredentialsHandler.LoadUserRegisterCredentialsPage,
			RequiresAuthentication: false,
		},
		{
			URI:                    "/register/credentials",
			Method:                 http.MethodPost,
			Function:               userCredentialsRoutes.userCredentialsHandler.CreateUserCredentials,
			RequiresAuthentication: false,
		},
		{
			URI:                    "/login",
			Method:                 http.MethodPost,
			Function:               userCredentialsRoutes.userCredentialsHandler.Login,
			RequiresAuthentication: false,
		},
	}
}
