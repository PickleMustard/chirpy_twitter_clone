package middleware

import (
	"github.com/PickleMustard/chirpy_twitter_clone/internal/apiprocessing"
	"net/http"
)

func (cfg *ApiConfig) UserValidation() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, d *http.Request) {
		apiprocessing.ValidateUser(w, d, cfg.Database)
	})
}

func (cfg *ApiConfig) UserInformationUpdate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, d *http.Request) {
		apiprocessing.UpdateUserInformation(w, d, cfg.Database, cfg.JWT_Secret)
	})
}

func (cfg *ApiConfig) UserLogin() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, d *http.Request) {
		apiprocessing.LoginUser(w, d, cfg.Database, cfg.JWT_Secret)
	})
}
