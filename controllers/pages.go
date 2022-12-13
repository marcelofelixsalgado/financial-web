package controllers

import (
	"marcelofelixsalgado/financial-web/utils"
	"net/http"
)

func LoadLoginScreen(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

func LoadUserRegisterPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "register.html", nil)
}

func LoadUserRegisterCredentialsPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "register-credentials.html", nil)
}
