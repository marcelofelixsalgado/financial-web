package user

import (
	"log"
	"marcelofelixsalgado/financial-web/api/cookies"
	"marcelofelixsalgado/financial-web/api/responses"
	"marcelofelixsalgado/financial-web/api/responses/faults"
	"marcelofelixsalgado/financial-web/pkg/usecase/user/create"
	"marcelofelixsalgado/financial-web/pkg/usecase/user/delete"
	"marcelofelixsalgado/financial-web/pkg/usecase/user/find"
	"marcelofelixsalgado/financial-web/pkg/usecase/user/update"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IUserHandler interface {
	CreateUser(ctx echo.Context) error
	UpdateUser(ctx echo.Context) error
	GetProfile(ctx echo.Context) error
	LoadUserEditPage(ctx echo.Context) error
	DeleteUser(ctx echo.Context) error
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

func (userHandler *UserHandler) CreateUser(ctx echo.Context) error {

	input := create.InputCreateUserDto{
		Name:  ctx.FormValue("name"),
		Phone: ctx.FormValue("phone"),
		Email: ctx.FormValue("email"),
	}

	// Validating input parameters
	if responseMessage := ValidateCreateRequestBody(input).GetMessage(); responseMessage.ErrorCode != "" {
		log.Printf("Error validating the request body: %v", responseMessage.GetMessage())
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	// Calling use case
	output, faultMessage, httpStatusCode, err := userHandler.createUseCase.Execute(input, ctx)
	if err != nil {
		log.Printf("Error trying to create the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	// Return error response
	if httpStatusCode != http.StatusCreated {
		log.Printf("Internal error: %d %v", httpStatusCode, faultMessage)
		return ctx.JSON(httpStatusCode, faultMessage)
	}

	// Response ok
	return ctx.JSON(httpStatusCode, output)
}

func (userHandler *UserHandler) UpdateUser(ctx echo.Context) error {

	cookie, _ := cookies.Read(ctx)
	loggedUserID := cookie.UserID

	input := update.InputUpdateUserDto{
		Id:    loggedUserID,
		Name:  ctx.FormValue("name"),
		Phone: ctx.FormValue("phone"),
		Email: ctx.FormValue("email"),
	}

	// Validating input parameters
	if responseMessage := ValidateUpdateRequestBody(input).GetMessage(); responseMessage.ErrorCode != "" {
		log.Printf("Error validating the request body: %v", responseMessage.GetMessage())
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	// Calling use case
	output, faultMessage, httpStatusCode, err := userHandler.updateUseCase.Execute(input, ctx)
	if err != nil {
		log.Printf("Error trying to update the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	// Return error response
	if httpStatusCode != http.StatusOK {
		log.Printf("Internal error: %d %v", httpStatusCode, faultMessage)
		return ctx.JSON(httpStatusCode, faultMessage)
	}

	// Response ok
	return ctx.JSON(httpStatusCode, output)
}

func (userHandler *UserHandler) GetProfile(ctx echo.Context) error {

	cookie, _ := cookies.Read(ctx)
	loggedUserID := cookie.UserID

	input := find.InputFindUserDto{
		Id: loggedUserID,
	}

	// Calling use case
	output, faultMessage, httpStatusCode, err := userHandler.findUseCase.Execute(input, ctx)
	if err != nil {
		log.Printf("Error trying to find the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	// Return error response
	if httpStatusCode != http.StatusOK {
		log.Printf("Internal error: %d %v", httpStatusCode, faultMessage)
		return ctx.JSON(httpStatusCode, faultMessage)
	}

	// Response ok
	return ctx.Render(http.StatusOK, "profile.html", output)
}

func (userHandler *UserHandler) LoadUserEditPage(ctx echo.Context) error {

	cookie, _ := cookies.Read(ctx)
	loggedUserID := cookie.UserID

	input := find.InputFindUserDto{
		Id: loggedUserID,
	}

	// Calling use case
	output, faultMessage, httpStatusCode, err := userHandler.findUseCase.Execute(input, ctx)
	if err != nil {
		log.Printf("Error trying to find the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	// Return error response
	if httpStatusCode != http.StatusOK {
		log.Printf("Internal error: %d %v", httpStatusCode, faultMessage)
		return ctx.JSON(httpStatusCode, faultMessage)
	}

	// Response ok
	return ctx.Render(http.StatusOK, "user-edit.html", output)
}

func (userHandler *UserHandler) DeleteUser(ctx echo.Context) error {

	cookie, _ := cookies.Read(ctx)
	loggedUserID := cookie.UserID

	input := delete.InputDeleteUserDto{
		Id: loggedUserID,
	}

	// Calling use case
	output, faultMessage, httpStatusCode, err := userHandler.deleteUseCase.Execute(input, ctx)
	if err != nil {
		log.Printf("Error trying to delete the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	// Return error response
	if httpStatusCode != http.StatusNoContent {
		log.Printf("Internal error: %d %v", httpStatusCode, faultMessage)
		return ctx.JSON(httpStatusCode, faultMessage)
	}

	// Response ok
	return ctx.JSON(httpStatusCode, output)
}
