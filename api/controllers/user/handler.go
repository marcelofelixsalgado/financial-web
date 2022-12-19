package user

import (
	"log"
	"marcelofelixsalgado/financial-web/api/responses"
	"marcelofelixsalgado/financial-web/api/responses/faults"
	"marcelofelixsalgado/financial-web/pkg/usecase/user/create"
	"net/http"
)

type IUserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	createUseCase create.ICreateUseCase
}

func NewUserHandler(createUseCase create.ICreateUseCase) IUserHandler {
	return &UserHandler{
		createUseCase: createUseCase,
	}
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
	output, faultMessage, httpStatusCode, err := userHandler.createUseCase.Execute(input)
	if err != nil {
		log.Printf("Error trying to convert the output to response body: %v", err)
		responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError).Write(w)
		return
	}

	// Return error response
	if httpStatusCode != http.StatusCreated {
		responses.JSON(w, httpStatusCode, faultMessage)
		return
	}

	// Response ok
	responses.JSON(w, httpStatusCode, output)
}
