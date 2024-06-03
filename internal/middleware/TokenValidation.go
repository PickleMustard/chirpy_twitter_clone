package middleware

import (
	"internal/apiprocessing"
	"net/http"
)

func (cfg *ApiConfig) RefreshToken() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, d *http.Request) {
		apiprocessing.ReturnAuthenticationToken(w, d, cfg.Database)
	})
}

func (cfg *ApiConfig) RevokeToken() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, d *http.Request) {
		apiprocessing.RevokeAuthenticationToken(w, d, cfg.Database)
	})
}
