package middlewares

import (
	"marcelofelixsalgado/financial-web/api/cookies"
	"marcelofelixsalgado/financial-web/api/responses"
	"marcelofelixsalgado/financial-web/api/responses/faults"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if _, err := cookies.Read(ctx); err != nil {
			http.Redirect(ctx.Response(), ctx.Request(), "/login", 302)
			responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.NotAuthorized)
			return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
		}
		return next(ctx)
	}
}
