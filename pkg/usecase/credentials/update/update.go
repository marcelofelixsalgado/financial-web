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
	Execute(InputUpdateUserCredentialsDto, *http.Request) (OutputUpdateUserCredentialsDto, faults.IFaultMessage, int, error)
}

type UpdateUseCase struct {
}

func NewUpdateUseCase() IUpdateUseCase {
	return &UpdateUseCase{}
}

func (UpdateUseCase *UpdateUseCase) Execute(input InputUpdateUserCredentialsDto, request *http.Request) (OutputUpdateUserCredentialsDto, faults.IFaultMessage, int, error) {

	var outputUpdateUserCredentialsDto OutputUpdateUserCredentialsDto
	user, err := json.Marshal(input)
	if err != nil {
		return OutputUpdateUserCredentialsDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	url := fmt.Sprintf("%s/v1/users/%s/credentials", configs.UserApiURL, input.UserId)
	response, err := requests.MakeUpstreamRequest(request, http.MethodPut, url, user, true)
	if err != nil {
		return OutputUpdateUserCredentialsDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return OutputUpdateUserCredentialsDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	if response.StatusCode >= 400 {
		var faultMessage faults.FaultMessage
		err := json.Unmarshal(bodyBytes, &faultMessage)
		if err != nil {
			return OutputUpdateUserCredentialsDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
		}
		return OutputUpdateUserCredentialsDto{}, faultMessage, response.StatusCode, nil
	}

	if err := json.Unmarshal(bodyBytes, &outputUpdateUserCredentialsDto); err != nil {
		return OutputUpdateUserCredentialsDto{}, faults.FaultMessage{}, http.StatusInternalServerError, err
	}

	return outputUpdateUserCredentialsDto, faults.FaultMessage{}, response.StatusCode, nil
}
