package credentials

import (
	"marcelofelixsalgado/financial-web/api/responses"
	"marcelofelixsalgado/financial-web/api/responses/faults"
	"marcelofelixsalgado/financial-web/pkg/usecase/credentials/create"
	"marcelofelixsalgado/financial-web/pkg/usecase/credentials/update"
)

func ValidateCreateRequestBody(inputCreateUserCredentialsDto create.InputCreateUserCredentialsDto) *responses.ResponseMessage {
	responseMessage := responses.NewResponseMessage()

	if inputCreateUserCredentialsDto.UserId == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "user.id", "")
	}

	if inputCreateUserCredentialsDto.Password == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "password", "")
	}

	return responseMessage
}

func ValidateUpdateRequestBody(inputUpdateUserCredentialsDto update.InputUpdateUserCredentialsDto) *responses.ResponseMessage {
	responseMessage := responses.NewResponseMessage()
	if inputUpdateUserCredentialsDto.UserId == "" {
		return responses.NewResponseMessage().AddMessageByIssue(faults.MissingRequiredField, responses.PathParameter, "id", "")
	}

	if inputUpdateUserCredentialsDto.UserId == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "user.id", "")
	}

	if inputUpdateUserCredentialsDto.NewPassword == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "new_password", "")
	}

	if inputUpdateUserCredentialsDto.CurrentPassword == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "current_password", "")
	}

	return responseMessage
}
