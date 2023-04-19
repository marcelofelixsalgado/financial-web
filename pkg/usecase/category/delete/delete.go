package delete

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

type IDeleteCategoryUseCase interface {
	Execute(InputDeleteCategoryDto, echo.Context) (OutputDeleteCategoryDto, faults.IFaultMessage, int, error)
}

type DeleteCategoryUseCase struct {
}

func NewDeleteCategoryUseCase() IDeleteCategoryUseCase {
	return &DeleteCategoryUseCase{}
}

func (deleteCategoryUseCase *DeleteCategoryUseCase) Execute(input InputDeleteCategoryDto, ctx echo.Context) (OutputDeleteCategoryDto, faults.IFaultMessage, int, error) {

	var outputDeleteCategoryDto OutputDeleteCategoryDto
	category, err := json.Marshal(input)
	if err != nil {
		return OutputDeleteCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	url := fmt.Sprintf("%s/v1/categories/%s", settings.Config.CategoryApiURL, input.Id)
	response, err := requests.MakeUpstreamRequest(ctx, http.MethodDelete, url, category, true)
	if err != nil {
		return OutputDeleteCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return OutputDeleteCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	if response.StatusCode >= 400 {
		var faultMessage faults.FaultMessage
		err := json.Unmarshal(bodyBytes, &faultMessage)
		if err != nil {
			return OutputDeleteCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
		}
		return OutputDeleteCategoryDto{}, faultMessage, response.StatusCode, nil
	}

	return outputDeleteCategoryDto, faults.FaultMessage{}, response.StatusCode, nil
}
