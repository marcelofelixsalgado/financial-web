package create

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

type ICreateSubCategoryUseCase interface {
	Execute(InputCreateSubCategoryDto, echo.Context) (OutputCreateSubCategoryDto, faults.IFaultMessage, int, error)
}

type CreateSubCategoryUseCase struct {
}

func NewCreateSubCategoryUseCase() ICreateSubCategoryUseCase {
	return &CreateSubCategoryUseCase{}
}

func (createSubCategoryUseCase *CreateSubCategoryUseCase) Execute(input InputCreateSubCategoryDto, ctx echo.Context) (OutputCreateSubCategoryDto, faults.IFaultMessage, int, error) {

	var outputCreateSubCategoryDto OutputCreateSubCategoryDto
	subCategory, err := json.Marshal(input)
	if err != nil {
		return OutputCreateSubCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	url := fmt.Sprintf("%s/v1/subcategories", settings.Config.CategoryApiURL)
	response, err := requests.MakeUpstreamRequest(ctx, http.MethodPost, url, subCategory, true)
	if err != nil {
		return OutputCreateSubCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return OutputCreateSubCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	if response.StatusCode >= 400 {
		var faultMessage faults.FaultMessage
		err := json.Unmarshal(bodyBytes, &faultMessage)
		if err != nil {
			return OutputCreateSubCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
		}
		return OutputCreateSubCategoryDto{}, faultMessage, response.StatusCode, nil
	}

	if err := json.Unmarshal(bodyBytes, &outputCreateSubCategoryDto); err != nil {
		return OutputCreateSubCategoryDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	return outputCreateSubCategoryDto, faults.FaultMessage{}, response.StatusCode, nil
}
