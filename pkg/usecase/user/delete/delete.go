package delete

import (
	"encoding/json"
	"fmt"
	"io"
	"marcelofelixsalgado/financial-web/configs"
	"marcelofelixsalgado/financial-web/pkg/infrastructure/requests"
	"marcelofelixsalgado/financial-web/pkg/usecase/responses/faults"
	"net/http"
)

type IDeleteUseCase interface {
	Execute(InputDeleteUserDto, *http.Request) (OutputDeleteUserDto, faults.IFaultMessage, int, error)
}

type DeleteUseCase struct {
}

func NewDeleteUseCase() IDeleteUseCase {
	return &DeleteUseCase{}
}

func (DeleteUseCase *DeleteUseCase) Execute(input InputDeleteUserDto, request *http.Request) (OutputDeleteUserDto, faults.IFaultMessage, int, error) {

	var outputDeleteUserDto OutputDeleteUserDto

	url := fmt.Sprintf("%s/v1/users/%s", configs.UserApiURL, input.Id)
	response, err := requests.MakeUpstreamRequest(request, http.MethodDelete, url, nil, true)
	if err != nil {
		return OutputDeleteUserDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			return OutputDeleteUserDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
		}

		var faultMessage faults.FaultMessage
		err = json.Unmarshal(bodyBytes, &faultMessage)
		if err != nil {
			return OutputDeleteUserDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
		}
		return OutputDeleteUserDto{}, faultMessage, response.StatusCode, nil
	}

	return outputDeleteUserDto, faults.FaultMessage{}, response.StatusCode, nil
}
