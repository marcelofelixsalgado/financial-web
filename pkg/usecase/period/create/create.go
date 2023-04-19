package create

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

type ICreatePeriodUseCase interface {
	Execute(InputCreatePeriodDto, echo.Context) (OutputCreatePeriodDto, faults.IFaultMessage, int, error)
}

type CreatePeriodUseCase struct {
}

func NewCreatePeriodUseCase() ICreatePeriodUseCase {
	return &CreatePeriodUseCase{}
}

func (createPeriodUseCase *CreatePeriodUseCase) Execute(input InputCreatePeriodDto, ctx echo.Context) (OutputCreatePeriodDto, faults.IFaultMessage, int, error) {

	var outputCreatePeriodDto OutputCreatePeriodDto
	period, err := json.Marshal(input)
	if err != nil {
		return OutputCreatePeriodDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	url := fmt.Sprintf("%s/v1/periods", settings.Config.PeriodApiURL)
	response, err := requests.MakeUpstreamRequest(ctx, http.MethodPost, url, period, true)
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
