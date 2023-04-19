package list

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/marcelofelixsalgado/financial-web/pkg/infrastructure/requests"
	"github.com/marcelofelixsalgado/financial-web/pkg/usecase/responses/faults"
	"github.com/marcelofelixsalgado/financial-web/settings"

	"github.com/labstack/echo/v4"
)

type IListTransactionTypeUseCase interface {
	Execute(InputListTransactionTypeDto, echo.Context) (OutputListTransactionTypeDto, faults.IFaultMessage, int, error)
}

type ListTransactionTypeUseCase struct {
}

func NewListTransactionTypeUseCase() IListTransactionTypeUseCase {
	return &ListTransactionTypeUseCase{}
}

func (listTransactionTypeUseCase *ListTransactionTypeUseCase) Execute(input InputListTransactionTypeDto, ctx echo.Context) (OutputListTransactionTypeDto, faults.IFaultMessage, int, error) {

	var transationTypes []TransactionType

	transationType, err := json.Marshal(input)
	if err != nil {
		return OutputListTransactionTypeDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	url := fmt.Sprintf("%s/v1/transaction_types", settings.Config.CategoryApiURL)

	response, err := requests.MakeUpstreamRequest(ctx, http.MethodGet, url, transationType, true)
	if err != nil {
		return OutputListTransactionTypeDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return OutputListTransactionTypeDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	if response.StatusCode >= 400 {
		var faultMessage faults.FaultMessage
		err := json.Unmarshal(bodyBytes, &faultMessage)
		if err != nil {
			return OutputListTransactionTypeDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
		}
		return OutputListTransactionTypeDto{}, faultMessage, response.StatusCode, nil
	}

	if err := json.Unmarshal(bodyBytes, &transationTypes); err != nil {
		return OutputListTransactionTypeDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	outputListTransactionTypeDto := OutputListTransactionTypeDto{
		TransactionTypes: transationTypes,
	}

	return outputListTransactionTypeDto, faults.FaultMessage{}, response.StatusCode, nil
}
