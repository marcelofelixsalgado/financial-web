package credentials

import (
	"marcelofelixsalgado/financial-web/api/controllers"
	"net/http"
)

var credentialsBasePath = "/register"

type UserCredentialsRoutes struct {
	userCredentialsHandler IUserCredentialsHandler
}

func NewUserCredentialsRoutes(userCredentialsHandler IUserCredentialsHandler) UserCredentialsRoutes {
	return UserCredentialsRoutes{
		userCredentialsHandler: userCredentialsHandler,
	}
}

func (userCredentialsRoutes *UserCredentialsRoutes) UserCredentialsRouteMapping() (string, []controllers.Route) {

	return credentialsBasePath, []controllers.Route{
		{
			URI:                    "",
			Method:                 http.MethodGet,
			Function:               userCredentialsRoutes.userCredentialsHandler.LoadUserRegisterPage,
			RequiresAuthentication: false,
		},
		{
			URI:                    "/credentials",
			Method:                 http.MethodGet,
			Function:               userCredentialsRoutes.userCredentialsHandler.LoadUserRegisterCredentialsPage,
			RequiresAuthentication: false,
		},
		{
			URI:                    "/credentials",
			Method:                 http.MethodPost,
			Function:               userCredentialsRoutes.userCredentialsHandler.CreateUserCredentials,
			RequiresAuthentication: false,
		},
		{
			URI:                    "/credentials",
			Method:                 http.MethodPut,
			Function:               userCredentialsRoutes.userCredentialsHandler.UpdateUserCredentials,
			RequiresAuthentication: false,
		},
		{
			URI:                    "/credentials/credentials-edit",
			Method:                 http.MethodGet,
			Function:               userCredentialsRoutes.userCredentialsHandler.LoadUserCredentialsEditPage,
			RequiresAuthentication: false,
		},
	}
}
