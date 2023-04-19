package period

import (
	"net/http"

	"github.com/marcelofelixsalgado/financial-web/api/controllers"
)

var periodsBasepath = "/periods"

type PeriodRoutes struct {
	periodHandler IPeriodHandler
}

func NewPeriodRoutes(periodHandler IPeriodHandler) PeriodRoutes {
	return PeriodRoutes{
		periodHandler: periodHandler,
	}
}

func (periodRoutes *PeriodRoutes) PeriodRouteMapping() (string, []controllers.Route) {

	return periodsBasepath, []controllers.Route{
		{
			URI:                    "",
			Method:                 http.MethodGet,
			Function:               periodRoutes.periodHandler.ListPeriod,
			RequiresAuthentication: true,
		},
		{
			URI:                    "",
			Method:                 http.MethodPost,
			Function:               periodRoutes.periodHandler.CreatePeriod,
			RequiresAuthentication: true,
		},
		{
			URI:                    "/:id",
			Method:                 http.MethodGet,
			Function:               periodRoutes.periodHandler.FindPeriod,
			RequiresAuthentication: true,
		},
		{
			URI:                    "/:id",
			Method:                 http.MethodPut,
			Function:               periodRoutes.periodHandler.UpdatePeriod,
			RequiresAuthentication: true,
		},
		{
			URI:                    "/:id",
			Method:                 http.MethodDelete,
			Function:               periodRoutes.periodHandler.DeletePeriod,
			RequiresAuthentication: true,
		},
	}
}
