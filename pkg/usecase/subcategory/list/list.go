package list

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

type IListSubCategoryUseCase interface {
	Execute(InputListSubCategoryDto, echo.Context) (OutputListSubCategoryDto, faults.IFaultMessage, int, error)
}

type ListSubCategoryUseCase struct {
}

func NewListSubCategoryUseCase() IListSubCategoryUseCase {
	return &ListSubCategoryUseCase{}
}

func (listSubCategoryUseCase *ListSubCategoryUseCase) Execute(input InputListSubCategoryDto, ctx echo.Context) (OutputListSubCategoryDto, faults.IFaultMessage, int, error) {

	var subCategories []SubCategory

	subCategory, err := json.Marshal(input)
	if err != nil {
		return OutputListSubCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	url := fmt.Sprintf("%s/v1/subcategories", settings.Config.CategoryApiURL)

	response, err := requests.MakeUpstreamRequest(ctx, http.MethodGet, url, subCategory, true)
	if err != nil {
		return OutputListSubCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return OutputListSubCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	if response.StatusCode >= 400 {
		var faultMessage faults.FaultMessage
		err := json.Unmarshal(bodyBytes, &faultMessage)
		if err != nil {
			return OutputListSubCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
		}
		return OutputListSubCategoryDto{}, faultMessage, response.StatusCode, nil
	}

	if err := json.Unmarshal(bodyBytes, &subCategories); err != nil {
		return OutputListSubCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	outputListSubCategoryDto := OutputListSubCategoryDto{
		SubCategories: subCategories,
	}

	return outputListSubCategoryDto, faults.FaultMessage{}, response.StatusCode, nil
}
