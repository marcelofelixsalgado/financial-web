package home

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type IHomeHandler interface {
	Home(ctx echo.Context) error
}

type HomeHandler struct {
}

func NewHomeHandler() IHomeHandler {
	return &HomeHandler{}
}

func (homeHandler *HomeHandler) Home(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "home.html", nil)
}
