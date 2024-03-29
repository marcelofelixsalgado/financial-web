package logout

import (
	"net/http"

	"github.com/marcelofelixsalgado/financial-web/api/controllers"
)

var logoutBasepath = "/logout"

type LogoutRoutes struct {
	LogoutHandler ILogoutHandler
}

func NewLogoutRoutes(LogoutHandler ILogoutHandler) LogoutRoutes {
	return LogoutRoutes{
		LogoutHandler: LogoutHandler,
	}
}

func (LogoutRoutes *LogoutRoutes) LogoutRouteMapping() (string, []controllers.Route) {
	return logoutBasepath, []controllers.Route{
		{
			URI:                    "",
			Method:                 http.MethodGet,
			Function:               LogoutRoutes.LogoutHandler.Logout,
			RequiresAuthentication: true,
		},
	}
}
