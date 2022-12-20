package period

import (
	"marcelofelixsalgado/financial-web/api/controllers"
	"net/http"
)

type PeriodRoutes struct {
	periodHandler IPeriodHandler
}

func NewPeriodRoutes(periodHandler IPeriodHandler) PeriodRoutes {
	return PeriodRoutes{
		periodHandler: periodHandler,
	}
}

func (periodRoutes *PeriodRoutes) PeriodRouteMapping() []controllers.Route {

	return []controllers.Route{
		{
			URI:                    "/periods",
			Method:                 http.MethodGet,
			Function:               periodRoutes.periodHandler.ListPeriod,
			RequiresAuthentication: true,
		},
		{
			URI:                    "/periods",
			Method:                 http.MethodPost,
			Function:               periodRoutes.periodHandler.CreatePeriod,
			RequiresAuthentication: true,
		},
		{
			URI:                    "/periods/{id}",
			Method:                 http.MethodGet,
			Function:               periodRoutes.periodHandler.FindPeriod,
			RequiresAuthentication: true,
		},
		{
			URI:                    "/periods/{id}",
			Method:                 http.MethodPut,
			Function:               periodRoutes.periodHandler.UpdatePeriod,
			RequiresAuthentication: true,
		},
		{
			URI:                    "/periods/{id}",
			Method:                 http.MethodDelete,
			Function:               periodRoutes.periodHandler.DeletePeriod,
			RequiresAuthentication: true,
		},
	}
}
