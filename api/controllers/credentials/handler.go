package credentials

import (
	"encoding/json"
	"io"
	"log"
	"marcelofelixsalgado/financial-web/api/responses"
	"marcelofelixsalgado/financial-web/api/responses/faults"
	"marcelofelixsalgado/financial-web/pkg/usecase/credentials/create"
	"marcelofelixsalgado/financial-web/pkg/usecase/credentials/login"
	loginUsecase "marcelofelixsalgado/financial-web/pkg/usecase/credentials/login"
	"net/http"

	"github.com/gorilla/mux"
)

type IUserCredentialsHandler interface {
	CreateUserCredentials(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type UserCredentialsHandler struct {
	createUseCase create.ICreateUseCase
	loginUseCase  login.ILoginUseCase
}

func NewUserCredentialsHandler(createUseCase create.ICreateUseCase, loginUseCase login.ILoginUseCase) IUserCredentialsHandler {
	return &UserCredentialsHandler{
		createUseCase: createUseCase,
		loginUseCase:  loginUseCase,
	}
}

func (userCredentialsHandler *UserCredentialsHandler) CreateUserCredentials(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	Id := parameters["id"]

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error trying to read the request body: %v", err)
		responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "").Write(w)
		return
	}

	var input create.InputCreateUserCredentialsDto

	if err := json.Unmarshal([]byte(requestBody), &input); err != nil {
		log.Printf("Error trying to convert the input data: %v", err)
		responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "").Write(w)
		return
	}
	input.UserId = Id

	// Validating input parameters
	if responseMessage := ValidateCreateRequestBody(input).GetMessage(); responseMessage.ErrorCode != "" {
		responseMessage.Write(w)
		return
	}

	output, faultMessage, err := userCredentialsHandler.createUseCase.Execute(input)
	if err != nil {
		log.Printf("Error trying to create the entity: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}

	if faultMessage.GetMessage().HttpStatusCode != http.StatusOK {
		responses.JSON(w, faultMessage.GetMessage().HttpStatusCode, faultMessage)
	}

	responses.JSON(w, faultMessage.GetMessage().HttpStatusCode, output)

	// outputJSON, err := json.Marshal(output)
	// if err != nil {
	// 	log.Printf("Error trying to convert the output to response body: %v", err)
	// 	responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
	// 	return
	// }

	// w.WriteHeader(http.StatusCreated)
	// w.Write(outputJSON)
}

func (userCredentialsHandler *UserCredentialsHandler) Login(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	input := loginUsecase.InputUserLoginDto{
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
	output, faultMessage, httpStatusCode, err := userCredentialsHandler.loginUseCase.Execute(input)
	if err != nil {
		log.Printf("Error trying to convert the output to response body: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}

	// Return error response
	if httpStatusCode != http.StatusOK {
		responses.JSON(w, faultMessage.GetMessage().HttpStatusCode, faultMessage)
	}

	// Response ok
	responses.JSON(w, httpStatusCode, output)
}
