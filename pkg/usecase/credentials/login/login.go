package login

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"marcelofelixsalgado/financial-web/configs"
	"marcelofelixsalgado/financial-web/pkg/usecase/responses/faults"
	"net/http"
)

type ILoginUseCase interface {
	Execute(InputUserLoginDto) (OutputUserLoginDto, faults.IFaultMessage, int, error)
}

type LoginUseCase struct {
}

func NewLoginUseCase() ILoginUseCase {
	return &LoginUseCase{}
}

func (loginUseCase *LoginUseCase) Execute(input InputUserLoginDto) (OutputUserLoginDto, faults.IFaultMessage, int, error) {
	var outputUserLoginDto OutputUserLoginDto
	user, err := json.Marshal(input)
	if err != nil {
		return OutputUserLoginDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	url := fmt.Sprintf("%s/v1/login", configs.UserApiURL)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(user))
	if err != nil {
		return OutputUserLoginDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return OutputUserLoginDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	if response.StatusCode >= 400 {
		var faultMessage faults.IFaultMessage
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
