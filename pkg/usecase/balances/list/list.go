package list

import (
	"encoding/json"
	"fmt"
	"io"
	"marcelofelixsalgado/financial-web/pkg/infrastructure/requests"
	"marcelofelixsalgado/financial-web/pkg/usecase/responses/faults"
	"marcelofelixsalgado/financial-web/settings"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IListBalanceUseCase interface {
	Execute(InputListBalanceDto, echo.Context) (OutputListBalanceDto, faults.IFaultMessage, int, error)
}

type ListBalanceUseCase struct {
}

func NewListBalanceUseCase() IListBalanceUseCase {
	return &ListBalanceUseCase{}
}

func (listBalanceUseCase *ListBalanceUseCase) Execute(input InputListBalanceDto, ctx echo.Context) (OutputListBalanceDto, faults.IFaultMessage, int, error) {

	var balances []Balance

	balance, err := json.Marshal(input)
	if err != nil {
		return OutputListBalanceDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	url := fmt.Sprintf("%s/v1/balances", settings.Config.BalanceApiURL)

	response, err := requests.MakeUpstreamRequest(ctx, http.MethodGet, url, balance, true)
	if err != nil {
		return OutputListBalanceDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return OutputListBalanceDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	if response.StatusCode >= 400 {
		var faultMessage faults.FaultMessage
		err := json.Unmarshal(bodyBytes, &faultMessage)
		if err != nil {
			return OutputListBalanceDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
		}
		return OutputListBalanceDto{}, faultMessage, response.StatusCode, nil
	}

	if err := json.Unmarshal(bodyBytes, &balances); err != nil {
		return OutputListBalanceDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	outputListBalanceDto := OutputListBalanceDto{
		Balances: balances,
	}

	return outputListBalanceDto, faults.FaultMessage{}, response.StatusCode, nil
}
