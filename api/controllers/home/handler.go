package home

import (
	"marcelofelixsalgado/financial-web/api/utils"
	"net/http"
)

type IHomeHandler interface {
	Home(w http.ResponseWriter, r *http.Request)
}

type HomeHandler struct {
}

func NewHomeHandler() IHomeHandler {
	return &HomeHandler{}
}

func (homeHandler *HomeHandler) Home(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "home.html", nil)
}
