package list

import (
	"encoding/json"
	"fmt"
	"io"
	"marcelofelixsalgado/financial-web/configs"
	"marcelofelixsalgado/financial-web/pkg/infrastructure/requests"
	"marcelofelixsalgado/financial-web/pkg/usecase/responses/faults"
	"net/http"
)

type IListPeriodUseCase interface {
	Execute(InputListPeriodDto, *http.Request) (OutputListPeriodDto, faults.IFaultMessage, int, error)
}

type ListPeriodUseCase struct {
}

func NewListPeriodUseCase() IListPeriodUseCase {
	return &ListPeriodUseCase{}
}

func (ListPeriodUseCase *ListPeriodUseCase) Execute(input InputListPeriodDto, request *http.Request) (OutputListPeriodDto, faults.IFaultMessage, int, error) {

	var periods []Period

	period, err := json.Marshal(input)
	if err != nil {
		return OutputListPeriodDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	url := fmt.Sprintf("%s/v1/periods", configs.PeriodApiURL)

	response, err := requests.MakeUpstreamRequest(request, http.MethodGet, url, period, true)
	if err != nil {
		return OutputListPeriodDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return OutputListPeriodDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	if response.StatusCode >= 400 {
		var faultMessage faults.FaultMessage
		err := json.Unmarshal(bodyBytes, &faultMessage)
		if err != nil {
			return OutputListPeriodDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
		}
		return OutputListPeriodDto{}, faultMessage, response.StatusCode, nil
	}

	if err := json.Unmarshal(bodyBytes, &periods); err != nil {
		return OutputListPeriodDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	outputListPeriodDto := OutputListPeriodDto{
		Periods: periods,
	}

	return outputListPeriodDto, faults.FaultMessage{}, response.StatusCode, nil
}