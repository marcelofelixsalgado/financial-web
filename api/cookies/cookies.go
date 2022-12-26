package cookies

import (
	"marcelofelixsalgado/financial-web/configs"
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
)

var secureCookie *securecookie.SecureCookie

// Uses the environment variables to create a SecureCookie
func Configure() {
	secureCookie = securecookie.New(configs.HashKey, configs.BlockKey)
}

// Register authentication information
func Save(w http.ResponseWriter, userID, accessToken string) error {

	data := map[string]string{
		"id":    userID,
		"token": accessToken,
	}

	encodedData, err := secureCookie.Encode("data", data)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    encodedData,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

// Return data stored in the cookie
func Read(r *http.Request) (map[string]string, error) {
	// read the cookie
	cookie, err := r.Cookie("data")
	if err != nil {
		return nil, err
	}

	// decode the data from the cookie
	values := make(map[string]string)
	if err = secureCookie.Decode("data", cookie.Value, &values); err != nil {
		return nil, err
	}
	return values, nil
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
