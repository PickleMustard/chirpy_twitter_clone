package main

import (
	"fmt"
	"internal/apiprocessing"
	"internal/database"
	"internal/endpoints"
	"internal/middleware"
	"log"
	"net/http"
	"os"
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
    log.Println(path)
	conf := middleware.ApiConfig{0}
	log.Println(conf.FileserverHits)
	chirpdb, dberr := database.NewDB(path)
	if dberr != nil {
		log.Fatal(dberr)
	}

	serve_mux := http.NewServeMux()
	serve_mux.Handle("/app/*", conf.MiddlewareMetricsIncrementor(http.StripPrefix("/app", http.FileServer(http.Dir(filepath_root)))))
	serve_mux.HandleFunc("GET /api/healthz", endpoints.ReadinessEndpointHandler)
	serve_mux.HandleFunc("GET /admin/metrics", conf.MetricsEndpointHandler)
	serve_mux.Handle("POST /api/chirps", apiprocessing.ValidateChirp(chirpdb))
    serve_mux.Handle("GET /api/chirps", apiprocessing.ReturnChirp(chirpdb))
    serve_mux.Handle("GET /api/chirps/{id}", apiprocessing.ReturnSpecificChirp(chirpdb))
    serve_mux.Handle("POST /api/users", apiprocessing.UserValidation(chirpdb))
	serve_mux.HandleFunc("/api/reset", conf.MiddlewareMetricsReset)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: serve_mux,
	}
	log.Printf("Serving files from %s on port: %s\n", filepath_root, port)
	log.Fatal(srv.ListenAndServe())
}
