package find

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

type IFindCategoryUseCase interface {
	Execute(InputFindCategoryDto, echo.Context) (OutputFindCategoryDto, faults.IFaultMessage, int, error)
}

type FindCategoryUseCase struct {
}

func NewFindCategoryUseCase() IFindCategoryUseCase {
	return &FindCategoryUseCase{}
}

func (FindCategoryUseCase *FindCategoryUseCase) Execute(input InputFindCategoryDto, ctx echo.Context) (OutputFindCategoryDto, faults.IFaultMessage, int, error) {

	var outputFindCategoryDto OutputFindCategoryDto
	category, err := json.Marshal(input)
	if err != nil {
		return OutputFindCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	url := fmt.Sprintf("%s/v1/categories/%s", settings.Config.CategoryApiURL, input.Id)
	response, err := requests.MakeUpstreamRequest(ctx, http.MethodGet, url, category, true)
	if err != nil {
		return OutputFindCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return OutputFindCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	if response.StatusCode >= 400 {
		var faultMessage faults.FaultMessage
		err := json.Unmarshal(bodyBytes, &faultMessage)
		if err != nil {
			return OutputFindCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
		}
		return OutputFindCategoryDto{}, faultMessage, response.StatusCode, nil
	}

	if err := json.Unmarshal(bodyBytes, &outputFindCategoryDto); err != nil {
		return OutputFindCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	return outputFindCategoryDto, faults.FaultMessage{}, response.StatusCode, nil
}
