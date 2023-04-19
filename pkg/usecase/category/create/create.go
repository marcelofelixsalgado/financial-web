package create

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

type ICreateCategoryUseCase interface {
	Execute(InputCreateCategoryDto, echo.Context) (OutputCreateCategoryDto, faults.IFaultMessage, int, error)
}

type CreateCategoryUseCase struct {
}

func NewCreateCategoryUseCase() ICreateCategoryUseCase {
	return &CreateCategoryUseCase{}
}

func (createCategoryUseCase *CreateCategoryUseCase) Execute(input InputCreateCategoryDto, ctx echo.Context) (OutputCreateCategoryDto, faults.IFaultMessage, int, error) {

	var outputCreateCategoryDto OutputCreateCategoryDto
	category, err := json.Marshal(input)
	if err != nil {
		return OutputCreateCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	url := fmt.Sprintf("%s/v1/categories", settings.Config.CategoryApiURL)
	response, err := requests.MakeUpstreamRequest(ctx, http.MethodPost, url, category, true)
	if err != nil {
		return OutputCreateCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return OutputCreateCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	if response.StatusCode >= 400 {
		var faultMessage faults.FaultMessage
		err := json.Unmarshal(bodyBytes, &faultMessage)
		if err != nil {
			return OutputCreateCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
		}
		return OutputCreateCategoryDto{}, faultMessage, response.StatusCode, nil
	}

	if err := json.Unmarshal(bodyBytes, &outputCreateCategoryDto); err != nil {
		return OutputCreateCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	return outputCreateCategoryDto, faults.FaultMessage{}, response.StatusCode, nil
}
