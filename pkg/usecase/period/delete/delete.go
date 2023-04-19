package delete

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

type IDeletePeriodUseCase interface {
	Execute(InputDeletePeriodDto, echo.Context) (OutputDeletePeriodDto, faults.IFaultMessage, int, error)
}

type DeletePeriodUseCase struct {
}

func NewDeletePeriodUseCase() IDeletePeriodUseCase {
	return &DeletePeriodUseCase{}
}

func (deletePeriodUseCase *DeletePeriodUseCase) Execute(input InputDeletePeriodDto, ctx echo.Context) (OutputDeletePeriodDto, faults.IFaultMessage, int, error) {

	var outputDeletePeriodDto OutputDeletePeriodDto
	period, err := json.Marshal(input)
	if err != nil {
		return OutputDeletePeriodDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	url := fmt.Sprintf("%s/v1/periods/%s", settings.Config.PeriodApiURL, input.Id)
	response, err := requests.MakeUpstreamRequest(ctx, http.MethodDelete, url, period, true)
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
