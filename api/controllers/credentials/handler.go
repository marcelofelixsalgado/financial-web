package credentials

import (
	"marcelofelixsalgado/financial-web/api/cookies"
	"marcelofelixsalgado/financial-web/api/responses"
	"marcelofelixsalgado/financial-web/api/responses/faults"
	"marcelofelixsalgado/financial-web/commons/logger"
	"marcelofelixsalgado/financial-web/pkg/usecase/credentials/create"
	"marcelofelixsalgado/financial-web/pkg/usecase/credentials/login"
	"marcelofelixsalgado/financial-web/pkg/usecase/credentials/update"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IUserCredentialsHandler interface {
	LoadUserRegisterPage(ctx echo.Context) error
	LoadUserRegisterCredentialsPage(ctx echo.Context) error
	CreateUserCredentials(ctx echo.Context) error
	UpdateUserCredentials(ctx echo.Context) error
	LoadUserCredentialsEditPage(ctx echo.Context) error
}

type UserCredentialsHandler struct {
	createUseCase create.ICreateUseCase
	updateUseCase update.IUpdateUseCase
	loginUseCase  login.ILoginUseCase
}

func NewUserCredentialsHandler(createUseCase create.ICreateUseCase, updateUseCase update.IUpdateUseCase, loginUseCase login.ILoginUseCase) IUserCredentialsHandler {
	return &UserCredentialsHandler{
		createUseCase: createUseCase,
		updateUseCase: updateUseCase,
		loginUseCase:  loginUseCase,
	}
}

func (userCredentialsHandler *UserCredentialsHandler) LoadUserRegisterPage(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "register.html", nil)
}

func (userCredentialsHandler *UserCredentialsHandler) LoadUserRegisterCredentialsPage(ctx echo.Context) error {

	return ctx.Render(http.StatusOK, "user-credentials.html", struct {
		User_id string
		Email   string
	}{
		User_id: ctx.FormValue("user_id"),
		Email:   ctx.FormValue("email"),
	})
}

func (userCredentialsHandler *UserCredentialsHandler) CreateUserCredentials(ctx echo.Context) error {
	log := logger.GetLogger()

	input := create.InputCreateUserCredentialsDto{
		UserId:   ctx.FormValue("user_id"),
		Password: ctx.FormValue("password"),
	}

	// Validating input parameters
	if responseMessage := ValidateCreateRequestBody(input).GetMessage(); responseMessage.ErrorCode != "" {
		log.Warnf("Error validating the request body: %v", responseMessage.GetMessage())
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	// Calling use case
	output, faultMessage, httpStatusCode, err := userCredentialsHandler.createUseCase.Execute(input, ctx)
	if err != nil {
		log.Warnf("Error trying to convert the output to response body: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	// Return error response
	if httpStatusCode != http.StatusCreated {
		log.Errorf("Internal error: %d %v", httpStatusCode, faultMessage)
		return ctx.JSON(httpStatusCode, faultMessage)
	}

	// Response ok
	return ctx.JSON(httpStatusCode, output)
}

func (userCredentialsHandler *UserCredentialsHandler) UpdateUserCredentials(ctx echo.Context) error {
	log := logger.GetLogger()

	cookie, _ := cookies.Read(ctx)
	loggedUserID := cookie.UserID

	input := update.InputUpdateUserCredentialsDto{
		UserId:          loggedUserID,
		CurrentPassword: ctx.FormValue("current_password"),
		NewPassword:     ctx.FormValue("new_password"),
	}

	// Validating input parameters
	if responseMessage := ValidateUpdateRequestBody(input).GetMessage(); responseMessage.ErrorCode != "" {
		log.Warnf("Error validating the request body: %v", responseMessage.GetMessage())
		ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	// Calling use case
	output, faultMessage, httpStatusCode, err := userCredentialsHandler.updateUseCase.Execute(input, ctx)
	if err != nil {
		log.Errorf("Error trying to convert the output to response body: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	// Return error response
	if httpStatusCode != http.StatusOK {
		log.Errorf("Internal error: %d %v", httpStatusCode, faultMessage)
		ctx.JSON(httpStatusCode, faultMessage)
	}

	// Response ok
	return ctx.JSON(httpStatusCode, output)
}

func (userCredentialsHandler *UserCredentialsHandler) LoadUserCredentialsEditPage(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "user-credentials-edit.html", nil)
}
