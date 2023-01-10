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

type ICreateUseCase interface {
	Execute(InputCreateUserCredentialsDto, echo.Context) (OutputCreateUserCredentialsDto, faults.IFaultMessage, int, error)
}

type CreateUseCase struct {
}

func NewCreateUseCase() ICreateUseCase {
	return &CreateUseCase{}
}

func (createUseCase *CreateUseCase) Execute(input InputCreateUserCredentialsDto, ctx echo.Context) (OutputCreateUserCredentialsDto, faults.IFaultMessage, int, error) {

	var outputCreateUserCredentialsDto OutputCreateUserCredentialsDto
	user, err := json.Marshal(input)
	if err != nil {
		return OutputCreateUserCredentialsDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	url := fmt.Sprintf("%s/v1/users/%s/credentials", settings.Config.UserApiURL, input.UserId)
	response, err := requests.MakeUpstreamRequest(ctx, http.MethodPost, url, user, false)
	if err != nil {
		return OutputCreateUserCredentialsDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return OutputCreateUserCredentialsDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	if response.StatusCode >= 400 {
		var faultMessage faults.FaultMessage
		err := json.Unmarshal(bodyBytes, &faultMessage)
		if err != nil {
			return OutputCreateUserCredentialsDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
		}
		return OutputCreateUserCredentialsDto{}, faultMessage, response.StatusCode, nil
	}

	if err := json.Unmarshal(bodyBytes, &outputCreateUserCredentialsDto); err != nil {
		return OutputCreateUserCredentialsDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	return outputCreateUserCredentialsDto, faults.FaultMessage{}, response.StatusCode, nil
}
