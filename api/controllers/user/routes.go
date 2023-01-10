package user

import (
	"marcelofelixsalgado/financial-web/api/controllers"
	"net/http"
)

var userBasepath = "/users"

type UserRoutes struct {
	userHandler IUserHandler
}

func NewUserRoutes(userHandler IUserHandler) UserRoutes {
	return UserRoutes{
		userHandler: userHandler,
	}
}

func (userRoutes *UserRoutes) UserRouteMapping() (string, []controllers.Route) {
	return userBasepath, []controllers.Route{
		{
			URI:                    "",
			Method:                 http.MethodPost,
			Function:               userRoutes.userHandler.CreateUser,
			RequiresAuthentication: false,
		},
		{
			URI:                    "/profile",
			Method:                 http.MethodGet,
			Function:               userRoutes.userHandler.GetProfile,
			RequiresAuthentication: true,
		},
		{
			URI:                    "/user-edit",
			Method:                 http.MethodGet,
			Function:               userRoutes.userHandler.LoadUserEditPage,
			RequiresAuthentication: true,
		},
		{
			URI:                    "",
			Method:                 http.MethodPut,
			Function:               userRoutes.userHandler.UpdateUser,
			RequiresAuthentication: true,
		},
		{
			URI:                    "",
			Method:                 http.MethodDelete,
			Function:               userRoutes.userHandler.DeleteUser,
			RequiresAuthentication: true,
		},
	}
}
