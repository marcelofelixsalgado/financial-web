package login

import (
	"net/http"

	"github.com/marcelofelixsalgado/financial-web/api/cookies"
	"github.com/marcelofelixsalgado/financial-web/api/responses"
	"github.com/marcelofelixsalgado/financial-web/api/responses/faults"
	"github.com/marcelofelixsalgado/financial-web/commons/logger"
	"github.com/marcelofelixsalgado/financial-web/pkg/usecase/credentials/login"

	"github.com/labstack/echo/v4"
)

type ILoginHandler interface {
	Login(ctx echo.Context) error
	LoadLoginPage(ctx echo.Context) error
}

type LoginHandler struct {
	loginUseCase login.ILoginUseCase
}

func NewLoginHandler(loginUseCase login.ILoginUseCase) ILoginHandler {
	return &LoginHandler{
		loginUseCase: loginUseCase,
	}
}

func (LoginHandler *LoginHandler) LoadLoginPage(ctx echo.Context) error {

	// If the user is already logged, doesn't make sense load the login page again
	cookie, _ := cookies.Read(ctx)
	if cookie.Token != "" {
		http.Redirect(ctx.Response().Writer, ctx.Request(), "/home", http.StatusFound)
		return ctx.JSON(http.StatusFound, nil)
	}

	return ctx.Render(http.StatusOK, "login.html", nil)
}

func (LoginHandler *LoginHandler) Login(ctx echo.Context) error {
	log := logger.GetLogger()

	input := login.InputUserLoginDto{
		Email:    ctx.FormValue("email"),
		Password: ctx.FormValue("password"),
	}

	// Validating input parameters
	if responseMessage := ValidateLoginRequestBody(input).GetMessage(); responseMessage.ErrorCode != "" {
		log.Warnf("Error validating the request body: %v", responseMessage.GetMessage())
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	// Calling use case
	output, faultMessage, httpStatusCode, err := LoginHandler.loginUseCase.Execute(input, ctx)
	if err != nil {
		log.Errorf("Error trying to convert the output to response body: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	// Returning backend error response
	if httpStatusCode != http.StatusOK {
		log.Errorf("Internal error: %d %v", httpStatusCode, faultMessage)
		return ctx.JSON(httpStatusCode, faultMessage)
	}

	// Saving the cookie
	if err = cookies.Save(ctx, output.User.Id, output.AccessToken); err != nil {
		log.Errorf("Error saving the cookie: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	// Response ok
	return ctx.JSON(httpStatusCode, output)
}

func (LoginHandler *LoginHandler) LoadLoginEditPage(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "user-credentials-edit.html", nil)
}
