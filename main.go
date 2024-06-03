package main

import (
	"fmt"
	"internal/database"
	"internal/endpoints"
	"internal/middleware"
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
		Database:       chirpdb,
		JWT_Secret:     os.Getenv("JWT_SECRET"),
	}
	log.Println(conf.JWT_Secret)

	serve_mux := http.NewServeMux()
	serve_mux.Handle("/app/*", conf.MiddlewareMetricsIncrementor(http.StripPrefix("/app", http.FileServer(http.Dir(filepath_root)))))
	serve_mux.HandleFunc("GET /api/healthz", endpoints.ReadinessEndpointHandler)
	serve_mux.HandleFunc("GET /admin/metrics", conf.MetricsEndpointHandler)
	serve_mux.Handle("POST /api/chirps", conf.CreateChirp())
	serve_mux.Handle("GET /api/chirps", conf.ReturnChirp())
	serve_mux.Handle("GET /api/chirps/{id}", conf.ReturnSpecificChirp())
  serve_mux.Handle("DELETE /api/chirps/{id}", conf.DeleteSpecificChirp())
	serve_mux.Handle("POST /api/users", conf.UserValidation())
	serve_mux.Handle("PUT /api/users", conf.UserInformationUpdate())
	serve_mux.Handle("POST /api/login", conf.UserLogin())
	serve_mux.HandleFunc("/api/reset", conf.MiddlewareMetricsReset)
  serve_mux.Handle("POST /api/refresh", conf.RefreshToken())
  serve_mux.Handle("POST /api/revoke", conf.RevokeToken())
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: serve_mux,
	}
	log.Printf("Serving files from %s on port: %s\n", filepath_root, port)
	log.Fatal(srv.ListenAndServe())
}
