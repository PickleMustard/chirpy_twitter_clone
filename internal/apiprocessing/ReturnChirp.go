package apiprocessing

import (
	"log"
	"net/http"
	"strconv"

	chirpdb "github.com/PickleMustard/chirpy_twitter_clone/internal/database"
)

func ReturnChirp(w http.ResponseWriter, d *http.Request, db *chirpdb.DB) error {
	author_id := d.URL.Query().Get("author_id")
	sort_order := d.URL.Query().Get("sort")
	respBody, err := db.GetChirps()
	if err != nil {
		log.Printf("Bad chirp response")
		RespondWithError(w, 404, "No chirps in database")
		return err
	}
	if len(author_id) > 0 {
		auid, _ := strconv.Atoi(author_id)
		respBody = limit_chirps_to_author(auid, respBody)
	}
	if sort_order == "desc" {
		respBody = reverse_chirp_order(respBody)
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

func limit_chirps_to_author(author_id int, response_list []chirpdb.Chirp) []chirpdb.Chirp {
	matching_chirps := make([]chirpdb.Chirp, 0)
	for _, chirp := range response_list {
		if chirp.Author_ID == author_id {
			matching_chirps = append(matching_chirps, chirp)
		}
	}
	return matching_chirps
}

func reverse_chirp_order(response_list []chirpdb.Chirp) []chirpdb.Chirp {
	for i, j := 0, len(response_list)-1; i < j; i, j = i+1, j-1 {
		response_list[i], response_list[j] = response_list[j], response_list[i]
	}
	return response_list
}
