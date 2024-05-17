package middleware

import (
	"internal/database"
)

type ApiConfig struct {
	FileserverHits int
	Database       *database.DB
	JWT_Secret     string
}
