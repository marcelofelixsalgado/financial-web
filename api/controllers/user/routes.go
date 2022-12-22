package user

import (
	"marcelofelixsalgado/financial-web/api/controllers"
	"net/http"
)

type UserRoutes struct {
	userHandler IUserHandler
}

func NewUserRoutes(userHandler IUserHandler) UserRoutes {
	return UserRoutes{
		userHandler: userHandler,
	}
}

func (userRoutes *UserRoutes) UserRouteMapping() []controllers.Route {
	return []controllers.Route{
		{
			URI:                    "/register",
			Method:                 http.MethodGet,
			Function:               userRoutes.userHandler.LoadUserRegisterPage,
			RequiresAuthentication: false,
		},
		{
			URI:                    "/users",
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
			URI:                    "/users",
			Method:                 http.MethodPut,
			Function:               userRoutes.userHandler.UpdateUser,
			RequiresAuthentication: true,
		},
		{
			URI:                    "/users",
			Method:                 http.MethodDelete,
			Function:               userRoutes.userHandler.DeleteUser,
			RequiresAuthentication: true,
		},
	}
}
