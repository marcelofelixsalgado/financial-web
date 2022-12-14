package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"marcelofelixsalgado/financial-web/responses"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	login, err := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		log.Fatal(err)
	}

	response, err := http.Post("http://localhost:8081/v1/login", "application/json", bytes.NewBuffer(login))
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.FormatErrorStatusCode(w, response)
		return
	}

	token, _ := ioutil.ReadAll(response.Body)
	fmt.Println(token)

	responses.JSON(w, response.StatusCode, nil)
}
