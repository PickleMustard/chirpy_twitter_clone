package apiprocessing

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

type LoginClaim struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func CreateNewTokenWithClaim(email, id, signing_key string, expiration_time time.Duration) (string, error) {
	if expiration_time <= 0 {
		expiration_time = time.Duration(time.Hour * 24)
	}
	user_claim := LoginClaim{
		email,
		jwt.RegisteredClaims{
			Issuer:    "Chirpy",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration_time)),
			Subject:   id,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user_claim)
	ss, err := token.SignedString(signing_key)
	if err != nil {
		log.Printf("Could not create user login token: %s", err)
		return "", err
	}

	return ss, nil
}

func ParseTokenWithClaim(token_string, auth_string string) (string, error) {
	token, parse_error := jwt.ParseWithClaims(token_string, &LoginClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(auth_string), nil
	})
    var id string
	if parse_error != nil {
		log.Printf("Error parsing token: %s", parse_error)
		return "", parse_error
	} else if claims, ok := token.Claims.(*LoginClaim); ok {
        if claims.RegisteredClaims.ExpiresAt.Time.Before(time.Now()) {
            log.Printf("Expired token, cannot procede")
            return "", errors.New("Expired token, cannot procede")
        }
        id = claims.RegisteredClaims.Subject

	} else {
		log.Printf("unknown claims type, cannot procede")
		return "", errors.New("Unknown claims type, cannot procede")
	}

	return id, nil

}
