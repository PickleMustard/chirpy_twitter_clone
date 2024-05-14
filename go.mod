module github.com/PickleMustard/chirpy_twitter_clone

go 1.22.2

require internal/middleware v1.0.0

require internal/endpoints v1.0.0

require internal/apiprocessing v1.0.0

require internal/database v1.0.0 // indirect

require github.com/google/uuid v1.6.0 // indirect

replace internal/middleware => ./internal/middleware

replace internal/endpoints => ./internal/endpoints

replace internal/apiprocessing => ./internal/api_processing

replace internal/database => ./internal/database
