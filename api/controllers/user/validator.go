package user

import (
	"marcelofelixsalgado/financial-web/api/responses"
	"marcelofelixsalgado/financial-web/api/responses/faults"
	"marcelofelixsalgado/financial-web/pkg/usecase/user/create"
)

type InputUserDto struct {
	name  string
	phone string
	email string
}

func ValidateCreateRequestBody(inputCreateUserDto create.InputCreateUserDto) *responses.ResponseMessage {
	inputUserDto := InputUserDto{
		name:  inputCreateUserDto.Name,
		phone: inputCreateUserDto.Phone,
		email: inputCreateUserDto.Email,
	}
	return validateRequestBody(inputUserDto)
}

func validateRequestBody(inputUserDto InputUserDto) *responses.ResponseMessage {

	responseMessage := responses.NewResponseMessage()

	if inputUserDto.name == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "name", "")
	}

	if inputUserDto.phone == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "phone", "")
	}

	if inputUserDto.email == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "email", "")
	}
	return responseMessage
}