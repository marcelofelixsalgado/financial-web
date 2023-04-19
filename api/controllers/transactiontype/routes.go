package transactiontype

import (
	"marcelofelixsalgado/financial-web/api/controllers"
	"net/http"
)

var basepath = "/transaction_types"

type TransactionTypeRoutes struct {
	transactionTypeHandler ITransactionTypeHandler
}

func NewTransactionTypeRoutes(transactionTypeHandler ITransactionTypeHandler) TransactionTypeRoutes {
	return TransactionTypeRoutes{
		transactionTypeHandler: transactionTypeHandler,
	}
}

func (transactionTypeRoutes *TransactionTypeRoutes) TransactionTypeRouteMapping() (string, []controllers.Route) {

	return basepath, []controllers.Route{
		{
			URI:                    "",
			Method:                 http.MethodGet,
			Function:               transactionTypeRoutes.transactionTypeHandler.ListTransactionType,
			RequiresAuthentication: true,
		},
	}
}
