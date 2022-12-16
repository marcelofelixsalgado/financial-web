package create

import (
	"bytes"
	"encoding/json"
	"io"
	"marcelofelixsalgado/financial-web/pkg/usecase/responses/faults"
	"net/http"
)

type ICreateUseCase interface {
	Execute(InputCreateUserCredentialsDto) (OutputCreateUserCredentialsDto, faults.IFaultMessage, error)
}

type CreateUseCase struct {
}

func NewCreateUseCase() ICreateUseCase {
	return &CreateUseCase{}
}

func (createUseCase *CreateUseCase) Execute(input InputCreateUserCredentialsDto) (OutputCreateUserCredentialsDto, faults.IFaultMessage, error) {

	var outputCreateUserCredentialsDto OutputCreateUserCredentialsDto
	user, err := json.Marshal(input)
	if err != nil {
		return OutputCreateUserCredentialsDto{}, faults.FaultMessage{}, err
	}

	response, err := http.Post("http://localhost:8081/v1/users", "application/json", bytes.NewBuffer(user))
	if err != nil {
		return OutputCreateUserCredentialsDto{}, faults.FaultMessage{}, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return OutputCreateUserCredentialsDto{}, faults.FaultMessage{}, err
	}

	if response.StatusCode >= 400 {
		var faultMessage faults.IFaultMessage
		err := json.Unmarshal(bodyBytes, &faultMessage)
		if err != nil {
			return OutputCreateUserCredentialsDto{}, faults.FaultMessage{}, err
		}
		return OutputCreateUserCredentialsDto{}, faultMessage, nil
	}

	if err := json.Unmarshal(bodyBytes, &outputCreateUserCredentialsDto); err != nil {
		return OutputCreateUserCredentialsDto{}, faults.FaultMessage{}, err
	}

	return outputCreateUserCredentialsDto, faults.FaultMessage{}, nil
}
