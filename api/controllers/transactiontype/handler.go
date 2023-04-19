package transactiontype

import (
	"github.com/marcelofelixsalgado/financial-web/api/responses"
	"github.com/marcelofelixsalgado/financial-web/api/responses/faults"
	"github.com/marcelofelixsalgado/financial-web/commons/logger"
	"github.com/marcelofelixsalgado/financial-web/pkg/usecase/transactiontype/list"

	"net/http"

	"github.com/labstack/echo/v4"
)

type ITransactionTypeHandler interface {
	ListTransactionType(ctx echo.Context) error
}

type TransactionTypeHandler struct {
	listTransactionTypeUseCase list.IListTransactionTypeUseCase
}

func NewTransactionTypeHandler(listTransactionTypeUseCase list.IListTransactionTypeUseCase) ITransactionTypeHandler {
	return &TransactionTypeHandler{
		listTransactionTypeUseCase: listTransactionTypeUseCase,
	}
}

func (transactionTypeHandler *TransactionTypeHandler) ListTransactionType(ctx echo.Context) error {
	log := logger.GetLogger()

	input := list.InputListTransactionTypeDto{}

	// Calling use case
	output, faultMessage, httpStatusCode, err := transactionTypeHandler.listTransactionTypeUseCase.Execute(input, ctx)
	if err != nil {
		log.Errorf("Error trying to convert the output to response body: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	// Return error response
	if httpStatusCode != http.StatusOK && httpStatusCode != http.StatusNotFound {
		log.Errorf("Internal error: %d %v", httpStatusCode, faultMessage)
		return ctx.JSON(httpStatusCode, faultMessage)
	}

	// Response ok
	return ctx.Render(http.StatusOK, "transaction-type.html", struct {
		TransactionTypes []list.TransactionType
	}{
		TransactionTypes: output.TransactionTypes,
	})
}
