package balance

import (
	"marcelofelixsalgado/financial-web/api/controllers"
	"net/http"
)

var balancesBasepath = "/balances"

type BalanceRoutes struct {
	balanceHandler IBalanceHandler
}

func NewBalanceRoutes(balanceHandler IBalanceHandler) BalanceRoutes {
	return BalanceRoutes{
		balanceHandler: balanceHandler,
	}
}

func (balanceRoutes *BalanceRoutes) BalanceRouteMapping() (string, []controllers.Route) {

	return balancesBasepath, []controllers.Route{
		{
			URI:                    "/periods",
			Method:                 http.MethodGet,
			Function:               balanceRoutes.balanceHandler.ListPeriod,
			RequiresAuthentication: true,
		},
		{
			URI:                    "",
			Method:                 http.MethodGet,
			Function:               balanceRoutes.balanceHandler.ListBalance,
			RequiresAuthentication: true,
		},
	}
}
