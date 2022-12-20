package find

import (
	"encoding/json"
	"fmt"
	"io"
	"marcelofelixsalgado/financial-web/configs"
	"marcelofelixsalgado/financial-web/pkg/infrastructure/requests"
	"marcelofelixsalgado/financial-web/pkg/usecase/responses/faults"
	"net/http"
)

type IFindPeriodUseCase interface {
	Execute(InputFindPeriodDto, *http.Request) (OutputFindPeriodDto, faults.IFaultMessage, int, error)
}

type FindPeriodUseCase struct {
}

func NewFindPeriodUseCase() IFindPeriodUseCase {
	return &FindPeriodUseCase{}
}

func (FindPeriodUseCase *FindPeriodUseCase) Execute(input InputFindPeriodDto, request *http.Request) (OutputFindPeriodDto, faults.IFaultMessage, int, error) {

	var outputFindPeriodDto OutputFindPeriodDto
	period, err := json.Marshal(input)
	if err != nil {
		return OutputFindPeriodDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	url := fmt.Sprintf("%s/v1/periods/%s", configs.PeriodApiURL, input.Id)
	response, err := requests.MakeUpstreamRequest(request, http.MethodGet, url, period, true)
	if err != nil {
		return OutputFindPeriodDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return OutputFindPeriodDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	if response.StatusCode >= 400 {
		var faultMessage faults.FaultMessage
		err := json.Unmarshal(bodyBytes, &faultMessage)
		if err != nil {
			return OutputFindPeriodDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
		}
		return OutputFindPeriodDto{}, faultMessage, response.StatusCode, nil
	}

	if err := json.Unmarshal(bodyBytes, &outputFindPeriodDto); err != nil {
		return OutputFindPeriodDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	return outputFindPeriodDto, faults.FaultMessage{}, response.StatusCode, nil
}
