package apiprocessing

import (
	"net/http"
)

func RespondWithNoBody(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
}
