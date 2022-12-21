package health

import (
	"marcelofelixsalgado/financial-web/api/controllers"
	"net/http"
)

var healthBasepath = "/health"

type HealthRoutes struct {
	healthHandler IHealthHandler
}

func NewHealthRoutes(healthHandler IHealthHandler) HealthRoutes {
	return HealthRoutes{
		healthHandler: healthHandler,
	}
}

func (healthRoutes *HealthRoutes) HealthRouteMapping() []controllers.Route {
	return []controllers.Route{
		{
			URI:                    healthBasepath,
			Method:                 http.MethodGet,
			Function:               healthRoutes.healthHandler.Health,
			RequiresAuthentication: false,
		},
	}
}
