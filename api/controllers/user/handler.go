package user

import (
	"log"
	"marcelofelixsalgado/financial-web/api/cookies"
	"marcelofelixsalgado/financial-web/api/responses"
	"marcelofelixsalgado/financial-web/api/responses/faults"
	"marcelofelixsalgado/financial-web/api/utils"
	"marcelofelixsalgado/financial-web/pkg/usecase/user/create"
	"marcelofelixsalgado/financial-web/pkg/usecase/user/delete"
	"marcelofelixsalgado/financial-web/pkg/usecase/user/find"
	"marcelofelixsalgado/financial-web/pkg/usecase/user/update"
	"net/http"
)

type IUserHandler interface {
	LoadUserRegisterPage(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	GetProfile(w http.ResponseWriter, r *http.Request)
	LoadUserEditPage(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	createUseCase create.ICreateUseCase
	updateUseCase update.IUpdateUseCase
	findUseCase   find.IFindUseCase
	deleteUseCase delete.IDeleteUseCase
}

func NewUserHandler(createUseCase create.ICreateUseCase, updateUseCase update.IUpdateUseCase,
	findUseCase find.IFindUseCase, deleteUseCase delete.IDeleteUseCase) IUserHandler {
	return &UserHandler{
		createUseCase: createUseCase,
		updateUseCase: updateUseCase,
		findUseCase:   findUseCase,
		deleteUseCase: deleteUseCase,
	}
}

func (userHandler *UserHandler) LoadUserRegisterPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "register.html", nil)
}

func (userHandler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	input := create.InputCreateUserDto{
		Name:  r.FormValue("name"),
		Phone: r.FormValue("phone"),
		Email: r.FormValue("email"),
	}

	// Validating input parameters
	if responseMessage := ValidateCreateRequestBody(input).GetMessage(); responseMessage.ErrorCode != "" {
		log.Printf("Error validating the request body: %v", responseMessage.GetMessage())
		responseMessage.Write(w)
		return
	}

	// Calling use case
	output, faultMessage, httpStatusCode, err := userHandler.createUseCase.Execute(input, r)
	if err != nil {
		log.Printf("Error trying to create the entity: %v", err)
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

func (userHandler *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	cookie, _ := cookies.Read(r)
	loggedUserID, _ := cookie["id"]

	input := update.InputUpdateUserDto{
		Id:    loggedUserID,
		Name:  r.FormValue("name"),
		Phone: r.FormValue("phone"),
		Email: r.FormValue("email"),
	}

	// Validating input parameters
	if responseMessage := ValidateUpdateRequestBody(input).GetMessage(); responseMessage.ErrorCode != "" {
		log.Printf("Error validating the request body: %v", responseMessage.GetMessage())
		responseMessage.Write(w)
		return
	}

	// Calling use case
	output, faultMessage, httpStatusCode, err := userHandler.updateUseCase.Execute(input, r)
	if err != nil {
		log.Printf("Error trying to update the entity: %v", err)
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

func (userHandler *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	cookie, _ := cookies.Read(r)
	loggedUserID, _ := cookie["id"]

	input := find.InputFindUserDto{
		Id: loggedUserID,
	}

	// Calling use case
	output, faultMessage, httpStatusCode, err := userHandler.findUseCase.Execute(input, r)
	if err != nil {
		log.Printf("Error trying to find the entity: %v", err)
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
	utils.ExecuteTemplate(w, "profile.html", output)
}

func (userHandler *UserHandler) LoadUserEditPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	cookie, _ := cookies.Read(r)
	loggedUserID, _ := cookie["id"]

	input := find.InputFindUserDto{
		Id: loggedUserID,
	}

	// Calling use case
	output, faultMessage, httpStatusCode, err := userHandler.findUseCase.Execute(input, r)
	if err != nil {
		log.Printf("Error trying to find the entity: %v", err)
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
	utils.ExecuteTemplate(w, "user-edit.html", output)
}

func (userHandler *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	cookie, _ := cookies.Read(r)
	loggedUserID, _ := cookie["id"]

	input := delete.InputDeleteUserDto{
		Id: loggedUserID,
	}

	// Calling use case
	output, faultMessage, httpStatusCode, err := userHandler.deleteUseCase.Execute(input, r)
	if err != nil {
		log.Printf("Error trying to delete the entity: %v", err)
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
