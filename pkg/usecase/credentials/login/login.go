package login

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

type ILoginUseCase interface {
	Execute(InputUserLoginDto, echo.Context) (OutputUserLoginDto, faults.IFaultMessage, int, error)
}

type LoginUseCase struct {
}

func NewLoginUseCase() ILoginUseCase {
	return &LoginUseCase{}
}

func (loginUseCase *LoginUseCase) Execute(input InputUserLoginDto, ctx echo.Context) (OutputUserLoginDto, faults.IFaultMessage, int, error) {
	var outputUserLoginDto OutputUserLoginDto
	user, err := json.Marshal(input)
	if err != nil {
		return OutputUserLoginDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	url := fmt.Sprintf("%s/v1/login", settings.Config.UserApiURL)
	response, err := requests.MakeUpstreamRequest(ctx, http.MethodPost, url, user, false)
	if err != nil {
		return OutputUserLoginDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return OutputUserLoginDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	if response.StatusCode >= 400 {
		var faultMessage faults.FaultMessage
		err := json.Unmarshal(bodyBytes, &faultMessage)
		if err != nil {
			return OutputUserLoginDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
		}
		return OutputUserLoginDto{}, faultMessage, response.StatusCode, nil
	}

	if err := json.Unmarshal(bodyBytes, &outputUserLoginDto); err != nil {
		return OutputUserLoginDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	return outputUserLoginDto, faults.FaultMessage{}, response.StatusCode, nil
}
