package apiprocessing

import (
	"encoding/json"
	chirpdb "github.com/PickleMustard/chirpy_twitter_clone/internal/database"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func ValidateChirp(w http.ResponseWriter, d *http.Request, db *chirpdb.DB, auth_string string) error {
	_, unparsed_token, _ := strings.Cut(d.Header.Get("Authorization"), "Bearer ")
	id, parse_error := ParseTokenWithClaim(unparsed_token, auth_string)
	if parse_error != nil {
		log.Printf("Couldn't authorize token: %s", parse_error)
		RespondWithError(w, 401, "Unauthorized user")
		return parse_error
	}

	type chirp_validation_return_values struct {
		CreatedAt       time.Time `json:"created_at"`
		Id              int       `json:"id"`
		Author_Id       int       `json:"author_id"`
		ValidationError string    `json:"validation_error"`
		CleanedBody     string    `json:"body"`
	}

	decoder := json.NewDecoder(d.Body)
	unvalidated_chirp := chirpdb.Chirp{}
	err := decoder.Decode(&unvalidated_chirp)
	if err != nil {
		log.Printf("Error decoding the chirp: %s", err)
		w.WriteHeader(500)
		return err
	}
	log.Printf("Chirp %s of Length: %d", unvalidated_chirp.ChirpBody, len(unvalidated_chirp.ChirpBody))
	clean_output := clean_chirp(unvalidated_chirp.ChirpBody)
	documented_chirp, err := db.CreateChirp(clean_output, id)
	if err != nil {
		log.Printf("Error occured in db: %s", err)
	}

	if len(unvalidated_chirp.ChirpBody) > 140 {
		RespondWithError(w, 400, "Chirp is too long")
	} else {
		respBody := chirp_validation_return_values{
			CreatedAt:       time.Now(),
			Id:              documented_chirp.ID,
			Author_Id:       id,
			ValidationError: "none",
			CleanedBody:     clean_output,
		}

		RespondWithJSON(w, 201, respBody)
	}
  return nil
}

func clean_chirp(body string) string {
	var nono_word_map map[string]bool
	nono_word_map = make(map[string]bool)
	nono_word_map["kerfuffle"] = true
	nono_word_map["sharbert"] = true
	nono_word_map["fornax"] = true

	split_strings := strings.Split(body, " ")
	for i, str := range split_strings {
		_, ok := nono_word_map[strings.ToLower(str)]
		if ok {
			split_strings[i] = "****"
		}
	}

	return strings.Join(split_strings, " ")
}

func DeleteChirp(w http.ResponseWriter, d *http.Request, db *chirpdb.DB, auth_string string) error {
	_, unparsed_token, _ := strings.Cut(d.Header.Get("Authorization"), "Bearer ")
	id, parse_error := ParseTokenWithClaim(unparsed_token, auth_string)
	if parse_error != nil {
		log.Printf("Couldn't authorize token: %s", parse_error)
		RespondWithError(w, 403, "Unauthorized user")
		return parse_error
	}
	desired_id := d.PathValue("id")
	log.Printf("Desired ID: %s\n", desired_id)
	chirp_id, conversion_error := strconv.Atoi(desired_id)
	if conversion_error != nil {
		log.Printf("Could not convert")
		RespondWithError(w, 404, "Could not find that chirp")
		return conversion_error
	}
	err := db.DeleteChirp(chirp_id, id)
	if err != nil {
		log.Printf("Bad chirp response")
		RespondWithError(w, 403, "Could not find that chirp")
		return err
	}
	RespondWithNoBody(w, 204)
  return nil
}
