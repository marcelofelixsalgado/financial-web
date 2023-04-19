package find

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

type IFindSubCategoryUseCase interface {
	Execute(InputFindSubCategoryDto, echo.Context) (OutputFindSubCategoryDto, faults.IFaultMessage, int, error)
}

type FindSubCategoryUseCase struct {
}

func NewFindSubCategoryUseCase() IFindSubCategoryUseCase {
	return &FindSubCategoryUseCase{}
}

func (FindSubCategoryUseCase *FindSubCategoryUseCase) Execute(input InputFindSubCategoryDto, ctx echo.Context) (OutputFindSubCategoryDto, faults.IFaultMessage, int, error) {

	var outputFindSubCategoryDto OutputFindSubCategoryDto
	subCategory, err := json.Marshal(input)
	if err != nil {
		return OutputFindSubCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	url := fmt.Sprintf("%s/v1/subcategories/%s", settings.Config.CategoryApiURL, input.Id)
	response, err := requests.MakeUpstreamRequest(ctx, http.MethodGet, url, subCategory, true)
	if err != nil {
		return OutputFindSubCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return OutputFindSubCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	if response.StatusCode >= 400 {
		var faultMessage faults.FaultMessage
		err := json.Unmarshal(bodyBytes, &faultMessage)
		if err != nil {
			return OutputFindSubCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
		}
		return OutputFindSubCategoryDto{}, faultMessage, response.StatusCode, nil
	}

	if err := json.Unmarshal(bodyBytes, &outputFindSubCategoryDto); err != nil {
		return OutputFindSubCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	return outputFindSubCategoryDto, faults.FaultMessage{}, response.StatusCode, nil
}
