package middleware

import (
	"github.com/PickleMustard/chirpy_twitter_clone/internal/apiprocessing"
	"net/http"
)

func (cfg *ApiConfig) CreateChirp() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, d *http.Request) {
		apiprocessing.ValidateChirp(w, d, cfg.Database, cfg.JWT_Secret)
	})
}

func (cfg *ApiConfig) ReturnChirp() http.Handler {
  return http.HandlerFunc(func (w http.ResponseWriter, d *http.Request) {
    apiprocessing.ReturnChirp(w, d, cfg.Database)
  })
}

func (cfg *ApiConfig) ReturnSpecificChirp() http.Handler {
  return http.HandlerFunc(func (w http.ResponseWriter, d *http.Request) {
    apiprocessing.ReturnSpecificChirp(w, d, cfg.Database)
  })
}

func (cfg *ApiConfig) DeleteSpecificChirp() http.Handler {
  return http.HandlerFunc(func (w http.ResponseWriter, d *http.Request) {
    apiprocessing.DeleteChirp(w, d, cfg.Database, cfg.JWT_Secret)
  })
}
