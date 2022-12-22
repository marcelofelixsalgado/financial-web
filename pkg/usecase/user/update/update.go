package update

import (
	"encoding/json"
	"fmt"
	"io"
	"marcelofelixsalgado/financial-web/configs"
	"marcelofelixsalgado/financial-web/pkg/infrastructure/requests"
	"marcelofelixsalgado/financial-web/pkg/usecase/responses/faults"
	"net/http"
)

type IUpdateUseCase interface {
	Execute(InputUpdateUserDto, *http.Request) (OutputUpdateUserDto, faults.IFaultMessage, int, error)
}

type UpdateUseCase struct {
}

func NewUpdateUseCase() IUpdateUseCase {
	return &UpdateUseCase{}
}

func (UpdateUseCase *UpdateUseCase) Execute(input InputUpdateUserDto, request *http.Request) (OutputUpdateUserDto, faults.IFaultMessage, int, error) {

	var outputUpdateUserDto OutputUpdateUserDto
	user, err := json.Marshal(input)
	if err != nil {
		return OutputUpdateUserDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	url := fmt.Sprintf("%s/v1/users/%s", configs.UserApiURL, input.Id)
	response, err := requests.MakeUpstreamRequest(request, http.MethodPut, url, user, true)
	if err != nil {
		return OutputUpdateUserDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return OutputUpdateUserDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	if response.StatusCode >= 400 {
		var faultMessage faults.FaultMessage
		err := json.Unmarshal(bodyBytes, &faultMessage)
		if err != nil {
			return OutputUpdateUserDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
		}
		return OutputUpdateUserDto{}, faultMessage, response.StatusCode, nil
	}

	if err := json.Unmarshal(bodyBytes, &outputUpdateUserDto); err != nil {
		return OutputUpdateUserDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	return outputUpdateUserDto, faults.FaultMessage{}, response.StatusCode, nil
}