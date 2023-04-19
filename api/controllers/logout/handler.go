package logout

import (
	"net/http"

	"github.com/marcelofelixsalgado/financial-web/api/cookies"

	"github.com/labstack/echo/v4"
)

type ILogoutHandler interface {
	Logout(ctx echo.Context) error
}

type LogoutHandler struct {
}

func NewLogoutHandler() ILogoutHandler {
	return &LogoutHandler{}
}

func (h *LogoutHandler) Logout(ctx echo.Context) error {

	// Removing the cookie
	cookies.Delete(ctx.Response().Writer)

	http.Redirect(ctx.Response().Writer, ctx.Request(), "/login", 302)

	return nil
}
