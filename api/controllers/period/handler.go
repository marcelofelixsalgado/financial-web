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

	"github.com/gorilla/mux"
)

type IPeriodHandler interface {
	CreatePeriod(w http.ResponseWriter, r *http.Request)
	ListPeriod(w http.ResponseWriter, r *http.Request)
	FindPeriod(w http.ResponseWriter, r *http.Request)
	UpdatePeriod(w http.ResponseWriter, r *http.Request)
	DeletePeriod(w http.ResponseWriter, r *http.Request)
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

func (periodHandler *PeriodHandler) CreatePeriod(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	year, err := strconv.Atoi(r.FormValue("year"))
	if err != nil {
		log.Printf("Error trying to convert the year in request body: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}

	startDateStr := fmt.Sprintf("%s%s", r.FormValue("start_date"), "T00:00:00Z")
	startDate, err := utils.ConvertStringToDateTime(startDateStr)
	if err != nil {
		log.Printf("Error trying to convert the StartDate at field: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}

	endDateStr := fmt.Sprintf("%s%s", r.FormValue("end_date"), "T23:59:59Z")
	endDate, err := utils.ConvertStringToDateTime(endDateStr)
	if err != nil {
		log.Printf("Error trying to convert the EndDate at field: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}

	input := create.InputCreatePeriodDto{
		Code:      r.FormValue("code"),
		Name:      r.FormValue("name"),
		Year:      year,
		StartDate: startDate,
		EndDate:   endDate,
	}

	// Validating input parameters
	if responseMessage := ValidateCreateRequestBody(input).GetMessage(); responseMessage.ErrorCode != "" {
		log.Printf("Error validating the request body: %v", responseMessage.GetMessage())
		responseMessage.Write(w)
		return
	}

	// Calling use case
	output, faultMessage, httpStatusCode, err := periodHandler.createPeriodUseCase.Execute(input, r)
	if err != nil {
		log.Printf("Error trying to convert the output to response body: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}

	// Return error response
	if httpStatusCode != http.StatusCreated {
		responses.JSON(w, httpStatusCode, faultMessage)
		log.Printf("Internal error: %d %v", httpStatusCode, faultMessage)
		return
	}

	// Response ok
	responses.JSON(w, httpStatusCode, output)
}

func (periodHandler *PeriodHandler) ListPeriod(w http.ResponseWriter, r *http.Request) {
	input := list.InputListPeriodDto{}

	// Calling use case
	output, faultMessage, httpStatusCode, err := periodHandler.listPeriodUseCase.Execute(input, r)
	if err != nil {
		log.Printf("Error trying to convert the output to response body: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}

	// Return error response
	if httpStatusCode != http.StatusOK && httpStatusCode != http.StatusNotFound {
		responses.JSON(w, httpStatusCode, faultMessage)
		log.Printf("Internal error: %d %v", httpStatusCode, faultMessage)
		return
	}

	// Response ok
	utils.ExecuteTemplate(w, "period.html", struct {
		Periods []list.Period
	}{
		Periods: output.Periods,
	})
}

func (periodHandler *PeriodHandler) FindPeriod(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	id := parameters["id"]

	input := find.InputFindPeriodDto{
		Id: id,
	}

	// Calling use case
	output, faultMessage, httpStatusCode, err := periodHandler.findPeriodUseCase.Execute(input, r)
	if err != nil {
		log.Printf("Error trying to convert the output to response body: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}

	// Return error response
	if httpStatusCode != http.StatusOK {
		responses.JSON(w, httpStatusCode, faultMessage)
		log.Printf("Internal error: %d %v", httpStatusCode, faultMessage)
		return
	}

	// Response ok
	utils.ExecuteTemplate(w, "period-edit.html", output)
}

func (periodHandler *PeriodHandler) UpdatePeriod(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	id := parameters["id"]

	year, err := strconv.Atoi(r.FormValue("year"))
	if err != nil {
		log.Printf("Error trying to convert the year in request body: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}

	startDateStr := fmt.Sprintf("%s%s", r.FormValue("start_date"), "T00:00:00Z")
	startDate, err := utils.ConvertStringToDateTime(startDateStr)
	if err != nil {
		log.Printf("Error trying to convert the StartDate at field: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}

	endDateStr := fmt.Sprintf("%s%s", r.FormValue("end_date"), "T23:59:59Z")
	endDate, err := utils.ConvertStringToDateTime(endDateStr)
	if err != nil {
		log.Printf("Error trying to convert the EndDate at field: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}

	input := update.InputUpdatePeriodDto{
		Id:        id,
		Code:      r.FormValue("code"),
		Name:      r.FormValue("name"),
		Year:      year,
		StartDate: startDate,
		EndDate:   endDate,
	}

	// Calling use case
	output, faultMessage, httpStatusCode, err := periodHandler.updatePeriodUseCase.Execute(input, r)
	if err != nil {
		log.Printf("Error trying to convert the output to response body: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}

	// Return error response
	if httpStatusCode != http.StatusOK {
		responses.JSON(w, httpStatusCode, faultMessage)
		log.Printf("Internal error: %d %v", httpStatusCode, faultMessage)
		return
	}

	// Response ok
	responses.JSON(w, httpStatusCode, output)
}

func (periodHandler *PeriodHandler) DeletePeriod(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	id := parameters["id"]

	input := delete.InputDeletePeriodDto{
		Id: id,
	}

	// Calling use case
	output, faultMessage, httpStatusCode, err := periodHandler.deletePeriodUseCase.Execute(input, r)
	if err != nil {
		log.Printf("Error trying to convert the output to response body: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}

	// Return error response
	if httpStatusCode != http.StatusNoContent {
		responses.JSON(w, httpStatusCode, faultMessage)
		log.Printf("Internal error: %d %v", httpStatusCode, faultMessage)
		return
	}

	// Response ok
	responses.JSON(w, httpStatusCode, output)
}
