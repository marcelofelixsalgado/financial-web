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

type IUpdateSubCategoryUseCase interface {
	Execute(InputUpdateSubCategoryDto, echo.Context) (OutputUpdateSubCategoryDto, faults.IFaultMessage, int, error)
}

type UpdateSubCategoryUseCase struct {
}

func NewUpdateSubCategoryUseCase() IUpdateSubCategoryUseCase {
	return &UpdateSubCategoryUseCase{}
}

func (UpdateSubCategoryUseCase *UpdateSubCategoryUseCase) Execute(input InputUpdateSubCategoryDto, ctx echo.Context) (OutputUpdateSubCategoryDto, faults.IFaultMessage, int, error) {

	var outputUpdateSubCategoryDto OutputUpdateSubCategoryDto
	subCategory, err := json.Marshal(input)
	if err != nil {
		return OutputUpdateSubCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	url := fmt.Sprintf("%s/v1/subcategories/%s", settings.Config.CategoryApiURL, input.Id)
	response, err := requests.MakeUpstreamRequest(ctx, http.MethodPut, url, subCategory, true)
	if err != nil {
		return OutputUpdateSubCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return OutputUpdateSubCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	if response.StatusCode >= 400 {
		var faultMessage faults.FaultMessage
		err := json.Unmarshal(bodyBytes, &faultMessage)
		if err != nil {
			return OutputUpdateSubCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
		}
		return OutputUpdateSubCategoryDto{}, faultMessage, response.StatusCode, nil
	}

	if err := json.Unmarshal(bodyBytes, &outputUpdateSubCategoryDto); err != nil {
		return OutputUpdateSubCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	return outputUpdateSubCategoryDto, faults.FaultMessage{}, response.StatusCode, nil
}
