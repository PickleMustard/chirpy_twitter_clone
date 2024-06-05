package apiprocessing

import (
	"encoding/json"
	"errors"
	chirpdb "github.com/PickleMustard/chirpy_twitter_clone/internal/database"
	"log"

	"net/http"
	"strings"
)

func UpgradeUser(w http.ResponseWriter, d *http.Request, db *chirpdb.DB, auth_string string) error {
	_, api_key, _ := strings.Cut(d.Header.Get("Authorization"), "ApiKey ")
  if api_key != auth_string {
    RespondWithError(w, 401, "")
    log.Printf("Unauthorized access of payment webhook attempted")
    return errors.New("Unauthorized access of payment webhook attempted")

  }
	  type payment_webhook struct {
    Event string `json:"event"`
    Data struct {
      User_id int `json:"user_id"`
    } `json:"data"`
  }
	decoder := json.NewDecoder(d.Body)
  unvalidated_webhook := payment_webhook{}
	err := decoder.Decode(&unvalidated_webhook)
	if err != nil {
		log.Printf("Error decoding the webhook: %s", err)
    RespondWithError(w, 500, "Error decoding the webhook")
		return err
	}

  if unvalidated_webhook.Event != "user.upgraded" {
    log.Printf("Don't care")
    RespondWithNoBody(w, 204)
    return nil
  }

  _, db_err := db.UpgradeUser(unvalidated_webhook.Data.User_id)
  if db_err != nil {
    log.Printf("Error upgrading user: %s", db_err)
    RespondWithError(w, 404, "Could not find user")
    return db_err
  }
  RespondWithNoBody(w, 204)
  return nil
}
