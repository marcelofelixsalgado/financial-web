package logout

import (
	"marcelofelixsalgado/financial-web/api/controllers"
	"net/http"
)

type LogoutRoutes struct {
	LogoutHandler ILogoutHandler
}

func NewLogoutRoutes(LogoutHandler ILogoutHandler) LogoutRoutes {
	return LogoutRoutes{
		LogoutHandler: LogoutHandler,
	}
}

func (LogoutRoutes *LogoutRoutes) LogoutRouteMapping() []controllers.Route {
	return []controllers.Route{
		{
			URI:                    "/logout",
			Method:                 http.MethodGet,
			Function:               LogoutRoutes.LogoutHandler.Logout,
			RequiresAuthentication: true,
		},
	}
}
