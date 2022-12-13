package controllers

import (
	"bytes"
	"encoding/json"
	"log"
	"marcelofelixsalgado/financial-web/responses"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name":  r.FormValue("name"),
		"phone": r.FormValue("phone"),
		"email": r.FormValue("email"),
	})
	if err != nil {
		log.Fatal(err)
	}

	response, err := http.Post("http://localhost:8081/v1/users", "application/json", bytes.NewBuffer(user))
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.FormatErrorStatusCode(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}
