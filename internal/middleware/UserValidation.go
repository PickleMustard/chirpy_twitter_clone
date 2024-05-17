package middleware

import (
	"internal/apiprocessing"
	"net/http"
)

func (cfg *ApiConfig) UserValidation() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, d *http.Request) {
		apiprocessing.ValidateUser(w, d, cfg.Database)
	})
}

func (cfg *ApiConfig) UserLogin() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, d *http.Request) {
		apiprocessing.LoginUser(w, d, cfg.Database, cfg.JWT_Secret)
	})
}
