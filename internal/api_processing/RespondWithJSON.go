package apiprocessing

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response_data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error in responding with JSON: %s", err)
		return
	}
	w.WriteHeader(code)
	w.Write(response_data)
}
