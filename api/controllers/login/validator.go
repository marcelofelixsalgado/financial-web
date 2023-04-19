package login

import (
	"github.com/marcelofelixsalgado/financial-web/api/responses"
	"github.com/marcelofelixsalgado/financial-web/api/responses/faults"
	"github.com/marcelofelixsalgado/financial-web/pkg/usecase/credentials/login"
)

func ValidateLoginRequestBody(inputUserLoginDto login.InputUserLoginDto) *responses.ResponseMessage {
	responseMessage := responses.NewResponseMessage()

	if inputUserLoginDto.Email == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "email", "")
	}

	if inputUserLoginDto.Password == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "password", "")
	}

	return responseMessage
}
