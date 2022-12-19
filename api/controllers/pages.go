package controllers

import (
	"marcelofelixsalgado/financial-web/utils"
	"net/http"
)

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

func LoadUserRegisterPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "register.html", nil)
}

func LoadUserRegisterCredentialsPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	utils.ExecuteTemplate(w, "registercredentials.html", struct {
		User_id string
	}{
		User_id: r.FormValue("user_id"),
	})
}

func LoadUserHomePage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "home.html", nil)
}
