package endpoints

import (
    "net/http"
)

func ReadinessEndpointHandler(w http.ResponseWriter, d *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
