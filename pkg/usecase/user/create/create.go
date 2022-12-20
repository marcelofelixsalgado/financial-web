package create

import (
	"encoding/json"
	"fmt"
	"io"
	"marcelofelixsalgado/financial-web/configs"
	"marcelofelixsalgado/financial-web/pkg/infrastructure/requests"
	"marcelofelixsalgado/financial-web/pkg/usecase/responses/faults"
	"net/http"
)

type ICreateUseCase interface {
	Execute(InputCreateUserDto, *http.Request) (OutputCreateUserDto, faults.IFaultMessage, int, error)
}

type CreateUseCase struct {
}

func NewCreateUseCase() ICreateUseCase {
	return &CreateUseCase{}
}

func (createUseCase *CreateUseCase) Execute(input InputCreateUserDto, request *http.Request) (OutputCreateUserDto, faults.IFaultMessage, int, error) {

	var outputCreateUserDto OutputCreateUserDto
	user, err := json.Marshal(input)
	if err != nil {
		return OutputCreateUserDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	url := fmt.Sprintf("%s/v1/users", configs.UserApiURL)
	response, err := requests.MakeUpstreamRequest(request, http.MethodPost, url, user, false)
	if err != nil {
		return OutputCreateUserDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return OutputCreateUserDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	if response.StatusCode >= 400 {
		var faultMessage faults.FaultMessage
		err := json.Unmarshal(bodyBytes, &faultMessage)
		if err != nil {
			return OutputCreateUserDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
		}
		return OutputCreateUserDto{}, faultMessage, response.StatusCode, nil
	}

	if err := json.Unmarshal(bodyBytes, &outputCreateUserDto); err != nil {
		return OutputCreateUserDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	return outputCreateUserDto, faults.FaultMessage{}, response.StatusCode, nil
}
