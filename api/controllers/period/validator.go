package period

import (
	"marcelofelixsalgado/financial-web/api/responses"
	"marcelofelixsalgado/financial-web/api/responses/faults"
	"marcelofelixsalgado/financial-web/pkg/usecase/period/create"
	"time"
)

type InputPeriodDto struct {
	code      string
	name      string
	year      int
	startDate time.Time
	endDate   time.Time
}

func ValidateCreateRequestBody(inputCreatePeriodDto create.InputCreatePeriodDto) *responses.ResponseMessage {
	inputPeriodDto := InputPeriodDto{
		code:      inputCreatePeriodDto.Code,
		name:      inputCreatePeriodDto.Name,
		year:      inputCreatePeriodDto.Year,
		startDate: inputCreatePeriodDto.StartDate,
		endDate:   inputCreatePeriodDto.EndDate,
	}
	return validateRequestBody(inputPeriodDto)
}

// func ValidateUpdateRequestBody(inputUpdatePeriodDto update.InputUpdatePeriodDto) *responses.ResponseMessage {
// 	if inputUpdatePeriodDto.Id == "" {
// 		return responses.NewResponseMessage().AddMessageByIssue(faults.MissingRequiredField, responses.PathParameter, "id", "")
// 	}
// 	inputPeriodDto := InputPeriodDto{
// 		code:      inputUpdatePeriodDto.Code,
// 		name:      inputUpdatePeriodDto.Name,
// 		year:      inputUpdatePeriodDto.Year,
// 		startDate: inputUpdatePeriodDto.StartDate,
// 		endDate:   inputUpdatePeriodDto.EndDate,
// 	}
// 	return validateRequestBody(inputPeriodDto)
// }

func validateRequestBody(inputPeriodDto InputPeriodDto) *responses.ResponseMessage {

	responseMessage := responses.NewResponseMessage()

	if inputPeriodDto.code == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "code", "")
	}

	if inputPeriodDto.name == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "name", "")
	}

	if inputPeriodDto.year == 0 {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "year", "")
	}

	if inputPeriodDto.startDate.IsZero() {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "start_date", "")
	}

	if inputPeriodDto.endDate.IsZero() {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "end_date", "")
	}

	if inputPeriodDto.startDate.Equal(inputPeriodDto.endDate) || inputPeriodDto.startDate.After(inputPeriodDto.endDate) {
		responseMessage.AddMessageByIssue(faults.ConditionalLowerThan, responses.Body, "start_date", inputPeriodDto.startDate.String(), "start_date", "end_date")
	}

	return responseMessage
}
