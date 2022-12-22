package create

import (
	"encoding/json"
	"fmt"
	"io"
	"marcelofelixsalgado/financial-web/configs"
	"marcelofelixsalgado/financial-web/pkg/infrastructure/requests"
	"marcelofelixsalgado/financial-web/pkg/usecase/responses/faults"
	"net/http"
)

type ICreatePeriodUseCase interface {
	Execute(InputCreatePeriodDto, *http.Request) (OutputCreatePeriodDto, faults.IFaultMessage, int, error)
}

type CreatePeriodUseCase struct {
}

func NewCreatePeriodUseCase() ICreatePeriodUseCase {
	return &CreatePeriodUseCase{}
}

func (createPeriodUseCase *CreatePeriodUseCase) Execute(input InputCreatePeriodDto, request *http.Request) (OutputCreatePeriodDto, faults.IFaultMessage, int, error) {

	var outputCreatePeriodDto OutputCreatePeriodDto
	period, err := json.Marshal(input)
	if err != nil {
		return OutputCreatePeriodDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	url := fmt.Sprintf("%s/v1/periods", configs.PeriodApiURL)
	response, err := requests.MakeUpstreamRequest(request, http.MethodPost, url, period, true)
	if err != nil {
		return OutputCreatePeriodDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return OutputCreatePeriodDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	if response.StatusCode >= 400 {
		var faultMessage faults.FaultMessage
		err := json.Unmarshal(bodyBytes, &faultMessage)
		if err != nil {
			return OutputCreatePeriodDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
		}
		return OutputCreatePeriodDto{}, faultMessage, response.StatusCode, nil
	}

	if err := json.Unmarshal(bodyBytes, &outputCreatePeriodDto); err != nil {
		return OutputCreatePeriodDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	return outputCreatePeriodDto, faults.FaultMessage{}, response.StatusCode, nil
}
