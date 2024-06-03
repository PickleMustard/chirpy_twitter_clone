package apiprocessing

import (
	chirpdb "internal/database"
	"log"
	"net/http"
	"strconv"
)

func ReturnChirp(db *chirpdb.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, d *http.Request) {
		respBody, err := db.GetChirps()
		if err != nil {
			log.Printf("Bad chirp response")
			RespondWithError(w, 404, "No chirps in database")
		}
		RespondWithJSON(w, 200, respBody)

	})
}

func ReturnSpecificChirp(db *chirpdb.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, d *http.Request) {
		desired_id := d.PathValue("id")
		log.Printf("Desired ID: %s\n", desired_id)
		id, conversion_error := strconv.Atoi(desired_id)
		if conversion_error != nil {
			log.Printf("Could not convert")
			RespondWithError(w, 404, "Could not find that chirp")
			return
		}
		respBody, err := db.GetSpecificChirps(id)
		if err != nil {
			log.Printf("Bad chirp response")
			RespondWithError(w, 404, "Could not find that chirp")
			return
		}
		RespondWithJSON(w, 200, respBody)

	})
}
