package create

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"marcelofelixsalgado/financial-web/configs"
	"marcelofelixsalgado/financial-web/pkg/usecase/responses/faults"
	"net/http"
)

type ICreateUseCase interface {
	Execute(InputCreateUserCredentialsDto) (OutputCreateUserCredentialsDto, faults.IFaultMessage, int, error)
}

type CreateUseCase struct {
}

func NewCreateUseCase() ICreateUseCase {
	return &CreateUseCase{}
}

func (createUseCase *CreateUseCase) Execute(input InputCreateUserCredentialsDto) (OutputCreateUserCredentialsDto, faults.IFaultMessage, int, error) {

	var outputCreateUserCredentialsDto OutputCreateUserCredentialsDto
	user, err := json.Marshal(input)
	if err != nil {
		return OutputCreateUserCredentialsDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	url := fmt.Sprintf("%s/v1/users/%s/credentials", configs.UserApiURL, input.UserId)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(user))
	if err != nil {
		return OutputCreateUserCredentialsDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return OutputCreateUserCredentialsDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	if response.StatusCode >= 400 {
		var faultMessage faults.IFaultMessage
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