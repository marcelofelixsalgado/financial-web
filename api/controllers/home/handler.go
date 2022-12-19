package home

import "net/http"

type IHomeHandler interface {
	Home(w http.ResponseWriter, r *http.Request)
}

type HomeHandler struct {
}

func NewHomeHandler() IHomeHandler {
	return &HomeHandler{}
}

func (homeHandler *HomeHandler) Home(w http.ResponseWriter, r *http.Request) {

}
