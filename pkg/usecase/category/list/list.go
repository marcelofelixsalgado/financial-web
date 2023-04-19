package list

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

type IListCategoryUseCase interface {
	Execute(InputListCategoryDto, echo.Context) (OutputListCategoryDto, faults.IFaultMessage, int, error)
}

type ListCategoryUseCase struct {
}

func NewListCategoryUseCase() IListCategoryUseCase {
	return &ListCategoryUseCase{}
}

func (listCategoryUseCase *ListCategoryUseCase) Execute(input InputListCategoryDto, ctx echo.Context) (OutputListCategoryDto, faults.IFaultMessage, int, error) {

	var categories []Category

	category, err := json.Marshal(input)
	if err != nil {
		return OutputListCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	url := fmt.Sprintf("%s/v1/categories", settings.Config.CategoryApiURL)

	response, err := requests.MakeUpstreamRequest(ctx, http.MethodGet, url, category, true)
	if err != nil {
		return OutputListCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return OutputListCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	if response.StatusCode >= 400 {
		var faultMessage faults.FaultMessage
		err := json.Unmarshal(bodyBytes, &faultMessage)
		if err != nil {
			return OutputListCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
		}
		return OutputListCategoryDto{}, faultMessage, response.StatusCode, nil
	}

	if err := json.Unmarshal(bodyBytes, &categories); err != nil {
		return OutputListCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	outputListCategoryDto := OutputListCategoryDto{
		Categories: categories,
	}

	return outputListCategoryDto, faults.FaultMessage{}, response.StatusCode, nil
}
