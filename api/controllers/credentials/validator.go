package credentials

import (
	"marcelofelixsalgado/financial-web/api/responses"
	"marcelofelixsalgado/financial-web/api/responses/faults"
	"marcelofelixsalgado/financial-web/pkg/usecase/credentials/create"
	"marcelofelixsalgado/financial-web/pkg/usecase/credentials/login"
)

type InputUserCredentialsDto struct {
	userId   string
	password string
}

func ValidateCreateRequestBody(inputCreateUserCredentialsDto create.InputCreateUserCredentialsDto) *responses.ResponseMessage {
	inputUserCredentialsDto := InputUserCredentialsDto{
		userId:   inputCreateUserCredentialsDto.UserId,
		password: inputCreateUserCredentialsDto.Password,
	}
	return validateRequestBody(inputUserCredentialsDto)
}

func validateRequestBody(inputUserCredentialsDto InputUserCredentialsDto) *responses.ResponseMessage {

	responseMessage := responses.NewResponseMessage()

	if inputUserCredentialsDto.userId == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "user.id", "")
	}

	if inputUserCredentialsDto.password == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "password", "")
	}

	return responseMessage
}

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
