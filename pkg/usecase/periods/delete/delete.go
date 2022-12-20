package delete

import (
	"encoding/json"
	"fmt"
	"io"
	"marcelofelixsalgado/financial-web/configs"
	"marcelofelixsalgado/financial-web/pkg/infrastructure/requests"
	"marcelofelixsalgado/financial-web/pkg/usecase/responses/faults"
	"net/http"
)

type IDeletePeriodUseCase interface {
	Execute(InputDeletePeriodDto, *http.Request) (OutputDeletePeriodDto, faults.IFaultMessage, int, error)
}

type DeletePeriodUseCase struct {
}

func NewDeletePeriodUseCase() IDeletePeriodUseCase {
	return &DeletePeriodUseCase{}
}

func (deletePeriodUseCase *DeletePeriodUseCase) Execute(input InputDeletePeriodDto, request *http.Request) (OutputDeletePeriodDto, faults.IFaultMessage, int, error) {

	var outputDeletePeriodDto OutputDeletePeriodDto
	period, err := json.Marshal(input)
	if err != nil {
		return OutputDeletePeriodDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	url := fmt.Sprintf("%s/v1/periods/%s", configs.PeriodApiURL, input.Id)
	response, err := requests.MakeUpstreamRequest(request, http.MethodDelete, url, period, true)
	if err != nil {
		return OutputDeletePeriodDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return OutputDeletePeriodDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	if response.StatusCode >= 400 {
		var faultMessage faults.FaultMessage
		err := json.Unmarshal(bodyBytes, &faultMessage)
		if err != nil {
			return OutputDeletePeriodDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
		}
		return OutputDeletePeriodDto{}, faultMessage, response.StatusCode, nil
	}

	return outputDeletePeriodDto, faults.FaultMessage{}, response.StatusCode, nil
}
