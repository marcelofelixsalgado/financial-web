package controllers

import (
	"marcelofelixsalgado/financial-web/utils"
	"net/http"
)

func LoadLoginScreen(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}
