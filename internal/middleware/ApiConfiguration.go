package middleware

import (
	"github.com/PickleMustard/chirpy_twitter_clone/internal/database"
)

type ApiConfig struct {
	FileserverHits int
	Database       *database.DB
	JWT_Secret     string
	Polka_Key      string
}
