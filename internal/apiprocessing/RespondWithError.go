package apiprocessing

import (
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(message))
}

func RespondWithInternalServerError(w http.ResponseWriter, code int) {
    w.WriteHeader(code)
}

