package period

import (
	"fmt"
	"log"
	"marcelofelixsalgado/financial-web/api/responses"
	"marcelofelixsalgado/financial-web/api/responses/faults"
	"marcelofelixsalgado/financial-web/api/utils"
	"marcelofelixsalgado/financial-web/pkg/usecase/periods/create"
	"marcelofelixsalgado/financial-web/pkg/usecase/periods/delete"
	"marcelofelixsalgado/financial-web/pkg/usecase/periods/find"
	"marcelofelixsalgado/financial-web/pkg/usecase/periods/list"
	"marcelofelixsalgado/financial-web/pkg/usecase/periods/update"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IPeriodHandler interface {
	CreatePeriod(ctx echo.Context) error
	ListPeriod(ctx echo.Context) error
	FindPeriod(ctx echo.Context) error
	UpdatePeriod(ctx echo.Context) error
	DeletePeriod(ctx echo.Context) error
}

type PeriodHandler struct {
	createPeriodUseCase create.ICreatePeriodUseCase
	listPeriodUseCase   list.IListPeriodUseCase
	findPeriodUseCase   find.IFindPeriodUseCase
	updatePeriodUseCase update.IUpdatePeriodUseCase
	deletePeriodUseCase delete.IDeletePeriodUseCase
}

func NewPeriodHandler(createPeriodUseCase create.ICreatePeriodUseCase, listPeriodUseCase list.IListPeriodUseCase,
	findPeriodUseCase find.IFindPeriodUseCase, updatePeriodUseCase update.IUpdatePeriodUseCase,
	periodDeleteUseCase delete.IDeletePeriodUseCase) IPeriodHandler {
	return &PeriodHandler{
		createPeriodUseCase: createPeriodUseCase,
		listPeriodUseCase:   listPeriodUseCase,
		findPeriodUseCase:   findPeriodUseCase,
		updatePeriodUseCase: updatePeriodUseCase,
		deletePeriodUseCase: periodDeleteUseCase,
	}
}

func (periodHandler *PeriodHandler) CreatePeriod(ctx echo.Context) error {

	year, err := strconv.Atoi(ctx.FormValue("year"))
	if err != nil {
		log.Printf("Error trying to convert the year in request body: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	startDateStr := fmt.Sprintf("%s%s", ctx.FormValue("start_date"), "T00:00:00Z")
	startDate, err := utils.ConvertStringToDateTime(startDateStr)
	if err != nil {
		log.Printf("Error trying to convert the StartDate at field: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	endDateStr := fmt.Sprintf("%s%s", ctx.FormValue("end_date"), "T23:59:59Z")
	endDate, err := utils.ConvertStringToDateTime(endDateStr)
	if err != nil {
		log.Printf("Error trying to convert the EndDate at field: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	input := create.InputCreatePeriodDto{
		Code:      ctx.FormValue("code"),
		Name:      ctx.FormValue("name"),
		Year:      year,
		StartDate: startDate,
		EndDate:   endDate,
	}

	// Validating input parameters
	if responseMessage := ValidateCreateRequestBody(input).GetMessage(); responseMessage.ErrorCode != "" {
		log.Printf("Error validating the request body: %v", responseMessage.GetMessage())
		ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	// Calling use case
	output, faultMessage, httpStatusCode, err := periodHandler.createPeriodUseCase.Execute(input, ctx)
	if err != nil {
		log.Printf("Error trying to convert the output to response body: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	// Return error response
	if httpStatusCode != http.StatusCreated {
		log.Printf("Internal error: %d %v", httpStatusCode, faultMessage)
		return ctx.JSON(httpStatusCode, faultMessage)
	}

	// Response ok
	return ctx.JSON(http.StatusOK, output)
}

func (periodHandler *PeriodHandler) ListPeriod(ctx echo.Context) error {
	input := list.InputListPeriodDto{}

	// Calling use case
	output, faultMessage, httpStatusCode, err := periodHandler.listPeriodUseCase.Execute(input, ctx)
	if err != nil {
		log.Printf("Error trying to convert the output to response body: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	// Return error response
	if httpStatusCode != http.StatusOK && httpStatusCode != http.StatusNotFound {
		log.Printf("Internal error: %d %v", httpStatusCode, faultMessage)
		return ctx.JSON(httpStatusCode, faultMessage)
	}

	// Response ok
	return ctx.Render(http.StatusOK, "period.html", struct {
		Periods []list.Period
	}{
		Periods: output.Periods,
	})
}

func (periodHandler *PeriodHandler) FindPeriod(ctx echo.Context) error {
	id := ctx.Param("id")

	input := find.InputFindPeriodDto{
		Id: id,
	}

	// Calling use case
	output, faultMessage, httpStatusCode, err := periodHandler.findPeriodUseCase.Execute(input, ctx)
	if err != nil {
		log.Printf("Error trying to convert the output to response body: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	// Return error response
	if httpStatusCode != http.StatusOK {
		log.Printf("Internal error: %d %v", httpStatusCode, faultMessage)
		return ctx.JSON(httpStatusCode, faultMessage)
	}

	// Response ok
	return ctx.Render(http.StatusOK, "period-edit.html", output)
}

func (periodHandler *PeriodHandler) UpdatePeriod(ctx echo.Context) error {
	id := ctx.Param("id")

	year, err := strconv.Atoi(ctx.FormValue("year"))
	if err != nil {
		log.Printf("Error trying to convert the year in request body: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	startDateStr := fmt.Sprintf("%s%s", ctx.FormValue("start_date"), "T00:00:00Z")
	startDate, err := utils.ConvertStringToDateTime(startDateStr)
	if err != nil {
		log.Printf("Error trying to convert the StartDate at field: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	endDateStr := fmt.Sprintf("%s%s", ctx.FormValue("end_date"), "T23:59:59Z")
	endDate, err := utils.ConvertStringToDateTime(endDateStr)
	if err != nil {
		log.Printf("Error trying to convert the EndDate at field: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	input := update.InputUpdatePeriodDto{
		Id:        id,
		Code:      ctx.FormValue("code"),
		Name:      ctx.FormValue("name"),
		Year:      year,
		StartDate: startDate,
		EndDate:   endDate,
	}

	// Calling use case
	output, faultMessage, httpStatusCode, err := periodHandler.updatePeriodUseCase.Execute(input, ctx)
	if err != nil {
		log.Printf("Error trying to convert the output to response body: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	// Return error response
	if httpStatusCode != http.StatusOK {
		log.Printf("Internal error: %d %v", httpStatusCode, faultMessage)
		return ctx.JSON(httpStatusCode, faultMessage)
	}

	// Response ok
	return ctx.JSON(http.StatusOK, output)
}

func (periodHandler *PeriodHandler) DeletePeriod(ctx echo.Context) error {
	id := ctx.Param("id")

	input := delete.InputDeletePeriodDto{
		Id: id,
	}

	// Calling use case
	output, faultMessage, httpStatusCode, err := periodHandler.deletePeriodUseCase.Execute(input, ctx)
	if err != nil {
		log.Printf("Error trying to convert the output to response body: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	// Return error response
	if httpStatusCode != http.StatusNoContent {
		log.Printf("Internal error: %d %v", httpStatusCode, faultMessage)
		return ctx.JSON(httpStatusCode, faultMessage)
	}

	// Response ok
	return ctx.JSON(http.StatusOK, output)
}
