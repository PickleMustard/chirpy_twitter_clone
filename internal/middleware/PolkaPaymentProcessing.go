package middleware

import (
  "net/http"
  "github.com/PickleMustard/chirpy_twitter_clone/internal/apiprocessing"
)

func (cfg *ApiConfig) PolkaUpgradeUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, d *http.Request) {
		apiprocessing.UpgradeUser(w, d, cfg.Database, cfg.Polka_Key)
	})
}
