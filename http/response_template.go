package http

import (
	"encoding/json"
	"log"
	"net/http"
)

// ResponseTemplate represents generic json response
type ResponseTemplate struct {
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

// EncodeError represents JSON response builder for error state
func EncodeError(w http.ResponseWriter, err error, code int, logger *log.Logger) {
	logger.Printf("http error: %s (code=%d)", err, code)

	// Write generic error response.
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(&ResponseTemplate{Message: "fail", Error: err.Error()})
}

// EncodeJSON represents a generic JSON response builder
func EncodeJSON(w http.ResponseWriter, v interface{}, logger *log.Logger) {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		EncodeError(w, err, http.StatusInternalServerError, logger)
	}
}
