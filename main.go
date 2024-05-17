package main

import (
	"fmt"
	"internal/apiprocessing"
    "internal/middleware"
	"internal/database"
	"internal/endpoints"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

//Handler Function for a readiness endpoint
//Matches the function signature of http.HandlerFunc

func main() {
	const filepath_root = "."
	const port = "8080"
    wd, err := os.Getwd()
    if err != nil {
        log.Fatal("Could not get directory")
    }
    path := fmt.Sprintf("%s/chirps.json", wd)
	chirpdb, dberr := database.NewDB(path)
	if dberr != nil {
		log.Fatal(dberr)
	}

    godotenv.Load()
    conf := middleware.ApiConfig{
        FileserverHits: 0,
        Database: chirpdb,
        JWT_Secret: os.Getenv("JWT_SECRET"),
    }
	log.Println(conf.JWT_Secret)

	serve_mux := http.NewServeMux()
	serve_mux.Handle("/app/*", conf.MiddlewareMetricsIncrementor(http.StripPrefix("/app", http.FileServer(http.Dir(filepath_root)))))
	serve_mux.HandleFunc("GET /api/healthz", endpoints.ReadinessEndpointHandler)
	serve_mux.HandleFunc("GET /admin/metrics", conf.MetricsEndpointHandler)
	serve_mux.Handle("POST /api/chirps", apiprocessing.ValidateChirp(chirpdb))
    serve_mux.Handle("GET /api/chirps", apiprocessing.ReturnChirp(chirpdb))
    serve_mux.Handle("GET /api/chirps/{id}", apiprocessing.ReturnSpecificChirp(chirpdb))
    serve_mux.Handle("POST /api/users", conf.UserValidation())
    serve_mux.Handle("POST /api/login", conf.UserLogin())
	serve_mux.HandleFunc("/api/reset", conf.MiddlewareMetricsReset)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: serve_mux,
	}
	log.Printf("Serving files from %s on port: %s\n", filepath_root, port)
	log.Fatal(srv.ListenAndServe())
}
