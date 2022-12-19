package requests

import (
	"fmt"
	"io"
	"marcelofelixsalgado/financial-web/api/cookies"
	"net/http"
)

// Make a request to another backend (upstream)
func MakeUpstreamRequest(r *http.Request, method, url string, data io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	// Get the access token from the cookie and set the upstream request header
	cookie, _ := cookies.Read(r)
	accessToken := fmt.Sprintf("Bearer %s", cookie["token"])
	request.Header.Add("Authorization", accessToken)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}