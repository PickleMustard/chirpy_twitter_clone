module github.com/PickleMustard/chirpy_twitter_clone

go 1.22.2

require internal/middleware v1.0.0
require internal/endpoints v1.0.0
require internal/apiprocessing v1.0.0

replace internal/middleware => ./internal/middleware
replace internal/endpoints => ./internal/endpoints
replace internal/apiprocessing => ./internal/api_processing
