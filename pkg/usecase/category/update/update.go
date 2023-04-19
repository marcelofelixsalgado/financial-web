package update

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

type IUpdateCategoryUseCase interface {
	Execute(InputUpdateCategoryDto, echo.Context) (OutputUpdateCategoryDto, faults.IFaultMessage, int, error)
}

type UpdateCategoryUseCase struct {
}

func NewUpdateCategoryUseCase() IUpdateCategoryUseCase {
	return &UpdateCategoryUseCase{}
}

func (UpdateCategoryUseCase *UpdateCategoryUseCase) Execute(input InputUpdateCategoryDto, ctx echo.Context) (OutputUpdateCategoryDto, faults.IFaultMessage, int, error) {

	var outputUpdateCategoryDto OutputUpdateCategoryDto
	category, err := json.Marshal(input)
	if err != nil {
		return OutputUpdateCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	url := fmt.Sprintf("%s/v1/categories/%s", settings.Config.CategoryApiURL, input.Id)
	response, err := requests.MakeUpstreamRequest(ctx, http.MethodPut, url, category, true)
	if err != nil {
		return OutputUpdateCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return OutputUpdateCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	if response.StatusCode >= 400 {
		var faultMessage faults.FaultMessage
		err := json.Unmarshal(bodyBytes, &faultMessage)
		if err != nil {
			return OutputUpdateCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
		}
		return OutputUpdateCategoryDto{}, faultMessage, response.StatusCode, nil
	}

	if err := json.Unmarshal(bodyBytes, &outputUpdateCategoryDto); err != nil {
		return OutputUpdateCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	return outputUpdateCategoryDto, faults.FaultMessage{}, response.StatusCode, nil
}
