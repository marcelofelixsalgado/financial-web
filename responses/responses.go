package responses

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type ResponseMessage struct {
	HttpStatusCode int                     `json:"-"`
	ErrorCode      string                  `json:"error_code"`
	Message        string                  `json:"message"`
	Details        []ResponseMessageDetail `json:"details,omitempty"`
}

type ResponseMessageDetail struct {
	Issue       string   `json:"issue"`
	Description string   `json:"description"`
	Location    Location `json:"location,omitempty"`
	Field       string   `json:"field,omitempty"`
	Value       string   `json:"value,omitempty"`
}

type Location string

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode != http.StatusNoContent {
		if erro := json.NewEncoder(w).Encode(data); erro != nil {
			log.Fatal(erro)
		}
	}
}

func FormatErrorStatusCode(w http.ResponseWriter, r *http.Response) {
	var responseMessage ResponseMessage

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Can not read body response")
		return
	}

	if err := json.Unmarshal(bodyBytes, &responseMessage); err != nil {
		log.Fatal("Can not unmarshal JSON")
		return
	}

	JSON(w, r.StatusCode, responseMessage)
}
