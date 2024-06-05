package apiprocessing

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
	ss, err := token.SignedString([]byte(signing_key))
	if err != nil {
		log.Printf("Could not create user login token: %s", err)
		return "", err
	}

	return ss, nil
}

func ParseTokenWithClaim(token_string, auth_string string) (int, error) {
	token, parse_error := jwt.ParseWithClaims(token_string, &LoginClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(auth_string), nil
	})
	var id string
	if parse_error != nil {
		log.Printf("Error parsing token: %s", parse_error)
		return 0, parse_error
	} else if claims, ok := token.Claims.(*LoginClaim); ok {
		if claims.RegisteredClaims.ExpiresAt.Time.Before(time.Now()) {
			log.Printf("Expired token, cannot procede")
			return 0, errors.New("Expired token, cannot procede")
		}
		id = claims.RegisteredClaims.Subject

	} else {
		log.Printf("unknown claims type, cannot procede")
		return 0, errors.New("Unknown claims type, cannot procede")
	}

	id_int, conv_err := strconv.Atoi(id)

	if conv_err != nil {
		log.Printf("Couldn't convert ID to int")
		return 0, conv_err
	}
	return id_int, nil
}
