package find

import (
	"encoding/json"
	"fmt"
	"io"
	"marcelofelixsalgado/financial-web/configs"
	"marcelofelixsalgado/financial-web/pkg/infrastructure/requests"
	"marcelofelixsalgado/financial-web/pkg/usecase/responses/faults"
	"net/http"
)

type IFindUseCase interface {
	Execute(InputFindUserDto, *http.Request) (OutputFindUserDto, faults.IFaultMessage, int, error)
}

type FindUseCase struct {
}

func NewFindUseCase() IFindUseCase {
	return &FindUseCase{}
}

func (FindUseCase *FindUseCase) Execute(input InputFindUserDto, request *http.Request) (OutputFindUserDto, faults.IFaultMessage, int, error) {

	var outputFindUserDto OutputFindUserDto
	user, err := json.Marshal(input)
	if err != nil {
		return OutputFindUserDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	url := fmt.Sprintf("%s/v1/users/%s", configs.UserApiURL, input.Id)
	response, err := requests.MakeUpstreamRequest(request, http.MethodGet, url, user, true)
	if err != nil {
		return OutputFindUserDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return OutputFindUserDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	if response.StatusCode >= 400 {
		var faultMessage faults.FaultMessage
		err := json.Unmarshal(bodyBytes, &faultMessage)
		if err != nil {
			return OutputFindUserDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
		}
		return OutputFindUserDto{}, faultMessage, response.StatusCode, nil
	}

	if err := json.Unmarshal(bodyBytes, &outputFindUserDto); err != nil {
		return OutputFindUserDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	return outputFindUserDto, faults.FaultMessage{}, response.StatusCode, nil
}