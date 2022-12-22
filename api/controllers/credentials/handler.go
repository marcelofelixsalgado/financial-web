package credentials

import (
	"log"
	"marcelofelixsalgado/financial-web/api/cookies"
	"marcelofelixsalgado/financial-web/api/responses"
	"marcelofelixsalgado/financial-web/api/responses/faults"
	"marcelofelixsalgado/financial-web/api/utils"
	"marcelofelixsalgado/financial-web/pkg/usecase/credentials/create"
	"marcelofelixsalgado/financial-web/pkg/usecase/credentials/login"
	"marcelofelixsalgado/financial-web/pkg/usecase/credentials/update"
	"net/http"
)

type IUserCredentialsHandler interface {
	LoadUserRegisterCredentialsPage(w http.ResponseWriter, r *http.Request)
	CreateUserCredentials(w http.ResponseWriter, r *http.Request)
	UpdateUserCredentials(w http.ResponseWriter, r *http.Request)
	LoadLoginPage(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	LoadUserCredentialsEditPage(w http.ResponseWriter, r *http.Request)
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

func (userCredentialsHandler *UserCredentialsHandler) LoadUserRegisterCredentialsPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	utils.ExecuteTemplate(w, "user-credentials.html", struct {
		User_id string
		Email   string
	}{
		User_id: r.FormValue("user_id"),
		Email:   r.FormValue("email"),
	})
}

func (userCredentialsHandler *UserCredentialsHandler) CreateUserCredentials(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	input := create.InputCreateUserCredentialsDto{
		UserId:   r.FormValue("user_id"),
		Password: r.FormValue("password"),
	}

	// Validating input parameters
	if responseMessage := ValidateCreateRequestBody(input).GetMessage(); responseMessage.ErrorCode != "" {
		log.Printf("Error validating the request body: %v", responseMessage.GetMessage())
		responseMessage.Write(w)
		return
	}

	// Calling use case
	output, faultMessage, httpStatusCode, err := userCredentialsHandler.createUseCase.Execute(input, r)
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

func (userCredentialsHandler *UserCredentialsHandler) UpdateUserCredentials(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	cookie, _ := cookies.Read(r)
	loggedUserID, _ := cookie["id"]

	input := update.InputUpdateUserCredentialsDto{
		UserId:          loggedUserID,
		CurrentPassword: r.FormValue("current_password"),
		NewPassword:     r.FormValue("new_password"),
	}

	// Validating input parameters
	if responseMessage := ValidateUpdateRequestBody(input).GetMessage(); responseMessage.ErrorCode != "" {
		log.Printf("Error validating the request body: %v", responseMessage.GetMessage())
		responseMessage.Write(w)
		return
	}

	// Calling use case
	output, faultMessage, httpStatusCode, err := userCredentialsHandler.updateUseCase.Execute(input, r)
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

func (userCredentialsHandler *UserCredentialsHandler) LoadLoginPage(w http.ResponseWriter, r *http.Request) {

	// If the user is already logged, doesn't make sense load the login page again
	cookie, _ := cookies.Read(r)
	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", 302)
		return
	}

	utils.ExecuteTemplate(w, "login.html", nil)
}

func (userCredentialsHandler *UserCredentialsHandler) Login(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	input := login.InputUserLoginDto{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	// Validating input parameters
	if responseMessage := ValidateLoginRequestBody(input).GetMessage(); responseMessage.ErrorCode != "" {
		log.Printf("Error validating the request body: %v", responseMessage.GetMessage())
		responseMessage.Write(w)
		return
	}

	// Calling use case
	output, faultMessage, httpStatusCode, err := userCredentialsHandler.loginUseCase.Execute(input, r)
	if err != nil {
		log.Printf("Error trying to convert the output to response body: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}

	// Returning backend error response
	if httpStatusCode != http.StatusOK {
		responses.JSON(w, httpStatusCode, faultMessage)
		log.Printf("Internal error: %d %v", httpStatusCode, faultMessage)
		return
	}

	// Saving the cookie
	if err = cookies.Save(w, output.User.Id, output.AccessToken); err != nil {
		log.Printf("Error saving the cookie: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}

	// Response ok
	responses.JSON(w, httpStatusCode, output)
}

func (userCredentialsHandler *UserCredentialsHandler) LoadUserCredentialsEditPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "user-credentials-edit.html", nil)
}
