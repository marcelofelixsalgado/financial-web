package cookies

import (
	"marcelofelixsalgado/financial-web/settings"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/meehow/securebytes"
)

type Session struct {
	UserID string
	Token  string
}

var secureCookie *securebytes.SecureBytes

// Uses the environment variables to create a SecureCookie
func Configure() {
	// secureCookie = securecookie.New("session", settings.Config.HashKey)
	secureCookie = securebytes.New(settings.Config.HashKey, securebytes.ASN1Serializer{})
}

// Register authentication information
func Save(ctx echo.Context, userID, accessToken string) error {

	session := Session{
		UserID: userID,
		Token:  accessToken,
	}

	encriptedData, err := secureCookie.EncryptToBase64(session)
	if err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name:     "data",
		Value:    encriptedData,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(ctx.Response(), cookie)

	// encodedData, err := secureCookie.Encode("data", data)
	// if err != nil {
	// 	return err
	// }

	// http.SetCookie(w, &http.Cookie{
	// 	Name:     "data",
	// 	Value:    encodedData,
	// 	Path:     "/",
	// 	HttpOnly: true,
	// })

	return nil
}

// Return data stored in the cookie
func Read(ctx echo.Context) (Session, error) {
	// read the cookie
	// cookie, err := ctx.Cookie("data")
	// if err != nil {
	// 	return nil, err
	// }

	// decode the data from the cookie
	// values := make(map[string]string)
	// if err = secureCookie.Decode("data", cookie.Value, &values); err != nil {
	// 	return nil, err
	// }

	var session Session
	cookie, err := ctx.Request().Cookie("data")
	if err != nil {
		return Session{}, err
	}
	if err = secureCookie.DecryptBase64(cookie.Value, &session); err != nil {
		return Session{}, err
	}

	return session, nil
}

// Remove cookie stored values
func Delete(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	})
}
