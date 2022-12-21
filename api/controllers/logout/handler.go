package logout

import (
	"marcelofelixsalgado/financial-web/api/cookies"
	"net/http"
)

type ILogoutHandler interface {
	Logout(w http.ResponseWriter, r *http.Request)
}

type LogoutHandler struct {
}

func NewLogoutHandler() ILogoutHandler {
	return &LogoutHandler{}
}

func (h *LogoutHandler) Logout(w http.ResponseWriter, r *http.Request) {

	// Removing the cookie
	cookies.Delete(w)

	http.Redirect(w, r, "/login", 302)
}
