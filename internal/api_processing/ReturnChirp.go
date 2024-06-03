package apiprocessing

import (
	chirpdb "internal/database"
	"log"
	"net/http"
	"strconv"
)

func ReturnChirp(w http.ResponseWriter, d *http.Request, db *chirpdb.DB) error {
	respBody, err := db.GetChirps()
	if err != nil {
		log.Printf("Bad chirp response")
		RespondWithError(w, 404, "No chirps in database")
		return err
	}
	RespondWithJSON(w, 200, respBody)
	return nil
}

func ReturnSpecificChirp(w http.ResponseWriter, d *http.Request, db *chirpdb.DB) error {
	desired_id := d.PathValue("id")
	log.Printf("Desired ID: %s\n", desired_id)
	id, conversion_error := strconv.Atoi(desired_id)
	if conversion_error != nil {
		log.Printf("Could not convert")
		RespondWithError(w, 404, "Could not find that chirp")
		return conversion_error
	}
	respBody, err := db.GetSpecificChirps(id)
	if err != nil {
		log.Printf("Bad chirp response")
		RespondWithError(w, 404, "Could not find that chirp")
		return err
	}
	RespondWithJSON(w, 200, respBody)
  return nil

}
