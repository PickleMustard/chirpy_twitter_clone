package apiprocessing

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	chirpdb "internal/database"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type UserInputValues struct {
	Email            string `json:"email"`
	Password         string `json:"password"`
	Expiration_Timer int    `json:"expires_in_seconds"`
}

type user_validation_return_values struct {
	CreatedAt       time.Time `json:"created_at"`
	Id              int       `json:"id"`
	ValidationError string    `json:"validation_error"`
	Body            string    `json:"email"`
}

type user_login_return_values struct {
	AttemptedLoginTime  time.Time `json:"login_attempt_time"`
	Id                  int       `json:"id"`
	AuthenticationError string    `json:"auth_error"`
	Email               string    `json:"email"`
	Token               string    `json:"token"`
	RefreshToken        string    `json:"refresh_token"`
}

type user_update_return_values struct {
	AttemptedLoginTime  time.Time `json:"login_attempt_time"`
	Id                  int       `json:"id"`
	AuthenticationError string    `json:"auth_error"`
	Email               string    `json:"email"`
}

func ValidateUser(w http.ResponseWriter, d *http.Request, db *chirpdb.DB) error {
	decoder := json.NewDecoder(d.Body)
	unvalidated_user := UserInputValues{}
	err := decoder.Decode(&unvalidated_user)
	if err != nil {
		log.Printf("Error getting user information: %s", err)
		w.WriteHeader(500)
		return err
	}

	if len(unvalidated_user.Email) <= 0 || len(unvalidated_user.Password) <= 0 {
		log.Printf("Either email or password required\n")
		w.WriteHeader(400)
		return errors.New("Either email or password required")
	}

	documented_user, err := db.CreateUser(unvalidated_user.Email, unvalidated_user.Password)
	if err != nil {
		log.Printf("Error occured in db: %s", err)
		return err
	}

	respBody := user_validation_return_values{
		CreatedAt:       time.Now(),
		Id:              documented_user.Id,
		ValidationError: "none",
		Body:            documented_user.Email,
	}

	RespondWithJSON(w, 201, respBody)
	return nil
}

func LoginUser(w http.ResponseWriter, d *http.Request, db *chirpdb.DB, auth_string string) error {
	decoder := json.NewDecoder(d.Body)
	unauthenticated_user := UserInputValues{}
	err := decoder.Decode(&unauthenticated_user)
	if err != nil {
		log.Printf("Error authenticating user: %s", err)
		w.WriteHeader(500)
		return err
	}

	if len(unauthenticated_user.Email) <= 0 || len(unauthenticated_user.Password) <= 0 {
		log.Printf("Either email or password required\n")
		w.WriteHeader(400)
		return errors.New("Either email or password required")
	}

	authenticated_user, auth_error := db.RetrieveUserByEmail(unauthenticated_user.Email, unauthenticated_user.Password)
	if auth_error != nil {
		log.Printf("Authentication Error Occurred: %s", auth_error)
		w.WriteHeader(401)
		return auth_error
	}
	token, err := CreateNewTokenWithClaim(authenticated_user.Email, strconv.Itoa(authenticated_user.Id), auth_string, time.Duration(unauthenticated_user.Expiration_Timer))
	if err != nil {
		log.Printf("Could not login user with token: %s", err)
		return err
	}
	refresh_token := make([]byte, 256)
	_, rand_err := rand.Read(refresh_token)

	if rand_err != nil {
		log.Printf("Error when making new refresh token: %s", rand_err)
		return rand_err
	}

	refresh_token_string := hex.EncodeToString(refresh_token)
  db.StoreToken(token,refresh_token_string)
	respBody := user_login_return_values{
		AttemptedLoginTime:  time.Now(),
		Id:                  authenticated_user.Id,
		AuthenticationError: "none",
		Email:               authenticated_user.Email,
		Token:               token,
		RefreshToken:        refresh_token_string,
	}

	RespondWithJSON(w, 200, respBody)
	return nil
}

func UpdateUserInformation(w http.ResponseWriter, d *http.Request, db *chirpdb.DB, auth_string string) error {
	_, unparsed_token, _ := strings.Cut(d.Header.Get("Authorization"), "Bearer ")
	id, parse_error := ParseTokenWithClaim(unparsed_token, auth_string)
	if parse_error != nil {
		log.Printf("Couldn't authorize token: %s", parse_error)
		RespondWithError(w, 401, "Unauthorized user")
		return parse_error
	}
	decoder := json.NewDecoder(d.Body)
	unauthenticated_user := UserInputValues{}
	err := decoder.Decode(&unauthenticated_user)
	if err != nil {
		log.Printf("Error decoding update values from request: %s", err)
		return err
	}

	updated_user, db_err := db.UpdateUser(unauthenticated_user.Email, unauthenticated_user.Password, id)
	if db_err != nil {
		log.Printf("Error updating user in the database: %s", db_err)
		return db_err
	}
	respBody := user_update_return_values{
		AttemptedLoginTime:  time.Now(),
		Id:                  id,
		AuthenticationError: "none",
		Email:               updated_user.Email,
	}

	RespondWithJSON(w, 200, respBody)
	return nil

}
