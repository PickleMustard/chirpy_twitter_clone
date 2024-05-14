package apiprocessing

import (
	"encoding/json"
	chirpdb "internal/database"
	"log"
	"net/http"
	"time"
)

func UserValidation(db *chirpdb.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, d *http.Request) {
		type user_validation_return_values struct {
			CreatedAt       time.Time `json:"created_at"`
			Id              int       `json:"id"`
			ValidationError string    `json:"validation_error"`
			Body            string    `json:"email"`
		}
		decoder := json.NewDecoder(d.Body)
		unvalidated_user := chirpdb.User{}
		err := decoder.Decode(&unvalidated_user)
		if err != nil {
			log.Printf("Error decoding the chirp: %s", err)
			w.WriteHeader(500)
			return
		}

		documented_user, err := db.CreateUser(unvalidated_user.Email)
		if err != nil {
			log.Printf("Error occured in db: %s", err)
		}

		respBody := user_validation_return_values{
			CreatedAt:       time.Now(),
			Id:              documented_user.Id,
			ValidationError: "none",
			Body:            documented_user.Email,
		}

		RespondWithJSON(w, 201, respBody)

	})
}
