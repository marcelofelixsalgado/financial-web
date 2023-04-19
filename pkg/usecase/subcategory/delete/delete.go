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

type IDeleteSubCategoryUseCase interface {
	Execute(InputDeleteSubCategoryDto, echo.Context) (OutputDeleteSubCategoryDto, faults.IFaultMessage, int, error)
}

type DeleteSubCategoryUseCase struct {
}

func NewDeleteSubCategoryUseCase() IDeleteSubCategoryUseCase {
	return &DeleteSubCategoryUseCase{}
}

func (deleteSubCategoryUseCase *DeleteSubCategoryUseCase) Execute(input InputDeleteSubCategoryDto, ctx echo.Context) (OutputDeleteSubCategoryDto, faults.IFaultMessage, int, error) {

	var outputDeleteSubCategoryDto OutputDeleteSubCategoryDto
	subCategory, err := json.Marshal(input)
	if err != nil {
		return OutputDeleteSubCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	url := fmt.Sprintf("%s/v1/subcategories/%s", settings.Config.CategoryApiURL, input.Id)
	response, err := requests.MakeUpstreamRequest(ctx, http.MethodDelete, url, subCategory, true)
	if err != nil {
		return OutputDeleteSubCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return OutputDeleteSubCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	if response.StatusCode >= 400 {
		var faultMessage faults.FaultMessage
		err := json.Unmarshal(bodyBytes, &faultMessage)
		if err != nil {
			return OutputDeleteSubCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
		}
		return OutputDeleteSubCategoryDto{}, faultMessage, response.StatusCode, nil
	}

	return outputDeleteSubCategoryDto, faults.FaultMessage{}, response.StatusCode, nil
}
