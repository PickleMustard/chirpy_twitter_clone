module github.com/PickleMustard/chirpy_twitter_clone

go 1.22.2

require internal/middleware v1.0.0

require internal/endpoints v1.0.0

require internal/apiprocessing v1.0.0

require internal/database v1.0.0 // indirect

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	golang.org/x/crypto v0.23.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/term v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
)

replace internal/middleware => ./internal/middleware

replace internal/endpoints => ./internal/endpoints

replace internal/apiprocessing => ./internal/api_processing

replace internal/database => ./internal/database
