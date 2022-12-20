package update

import (
	"encoding/json"
	"fmt"
	"io"
	"marcelofelixsalgado/financial-web/configs"
	"marcelofelixsalgado/financial-web/pkg/infrastructure/requests"
	"marcelofelixsalgado/financial-web/pkg/usecase/responses/faults"
	"net/http"
)

type IUpdatePeriodUseCase interface {
	Execute(InputUpdatePeriodDto, *http.Request) (OutputUpdatePeriodDto, faults.IFaultMessage, int, error)
}

type UpdatePeriodUseCase struct {
}

func NewUpdatePeriodUseCase() IUpdatePeriodUseCase {
	return &UpdatePeriodUseCase{}
}

func (UpdatePeriodUseCase *UpdatePeriodUseCase) Execute(input InputUpdatePeriodDto, request *http.Request) (OutputUpdatePeriodDto, faults.IFaultMessage, int, error) {

	var outputUpdatePeriodDto OutputUpdatePeriodDto
	period, err := json.Marshal(input)
	if err != nil {
		return OutputUpdatePeriodDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	url := fmt.Sprintf("%s/v1/periods/%s", configs.PeriodApiURL, input.Id)
	response, err := requests.MakeUpstreamRequest(request, http.MethodPut, url, period, true)
	if err != nil {
		return OutputUpdatePeriodDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return OutputUpdatePeriodDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	if response.StatusCode >= 400 {
		var faultMessage faults.FaultMessage
		err := json.Unmarshal(bodyBytes, &faultMessage)
		if err != nil {
			return OutputUpdatePeriodDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
		}
		return OutputUpdatePeriodDto{}, faultMessage, response.StatusCode, nil
	}

	if err := json.Unmarshal(bodyBytes, &outputUpdatePeriodDto); err != nil {
		return OutputUpdatePeriodDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	return outputUpdatePeriodDto, faults.FaultMessage{}, response.StatusCode, nil
}
