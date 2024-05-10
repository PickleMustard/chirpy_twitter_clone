package apiprocessing

import (
	"encoding/json"
	uuid "github.com/google/uuid"
	"log"
	"net/http"
	"strings"
	"time"
)

func ValidateChirp(w http.ResponseWriter, d *http.Request) {
	type chirp struct {
		ChirpBody string `json:"body"`
	}
	type chirp_validation_return_values struct {
		CreatedAt       time.Time `json:"created_at"`
		Id              string    `json:"id"`
		ValidationError string    `json:"validation_error"`
		CleanedBody     string    `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(d.Body)
	unvalidated_chirp := chirp{}
	err := decoder.Decode(&unvalidated_chirp)
	if err != nil {
		log.Printf("Error decoding the chirp: %s", err)
		w.WriteHeader(500)
		return
	}
	log.Printf("Chirp %s of Length: %d", unvalidated_chirp.ChirpBody, len(unvalidated_chirp.ChirpBody))
	clean_output := clean_chirp(unvalidated_chirp.ChirpBody)

	if len(unvalidated_chirp.ChirpBody) > 140 {
        RespondWithError(w, 400, "Chirp is too long")
	} else {
		respBody := chirp_validation_return_values{
			CreatedAt:       time.Now(),
			Id:              uuid.NewString(),
			ValidationError: "none",
			CleanedBody:     clean_output,
		}

        RespondWithJSON(w, 200, respBody)
	}
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
