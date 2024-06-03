package apiprocessing

import (
	"log"
	"net/http"
  "strings"
	chirpdb "internal/database"
)

type Refresh_Token_Return_Values struct {
  Auth_Token string `json:"token"`
}

func ReturnAuthenticationToken( w http.ResponseWriter, d *http.Request, db *chirpdb.DB) error {
	_, refresh_token, _ := strings.Cut(d.Header.Get("Authorization"), "Bearer ")

  log.Println(refresh_token)
 token, retrieval_err := db.RetrieveAuthToken(refresh_token)
  if retrieval_err != nil{
    RespondWithError(w, 401, "Authorization Error")
    log.Printf("Unable to retrieve auth token from the database: %s", retrieval_err)
    return retrieval_err
  }

  respBody := Refresh_Token_Return_Values{
    Auth_Token: token.Auth_Token,
  }

	RespondWithJSON(w, 200, respBody)
	return nil
}

func RevokeAuthenticationToken(w http.ResponseWriter, d *http.Request, db *chirpdb.DB) error {
  _, refresh_token, _ := strings.Cut(d.Header.Get("Authorization"), "Bearer ")

  deletion_err := db.DeleteAuthorizationToken(refresh_token)

  if deletion_err != nil {
    RespondWithError(w, 401, "Authorization Error")
    log.Printf("Error deleting auth token from the database: %s", deletion_err)
    return deletion_err
  }

  RespondWithNoBody(w, 204)
  return nil
}
