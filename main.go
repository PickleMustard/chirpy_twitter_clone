package main

import (
    "internal/apiprocessing"
    "internal/endpoints"
	"internal/middleware"
	"log"
	"net/http"
)

//Handler Function for a readiness endpoint
//Matches the function signature of http.HandlerFunc

func main() {
    const filepath_root = "."
    const port = "8080"
    conf := middleware.ApiConfig {0}
    log.Println(conf.FileserverHits)

    serve_mux := http.NewServeMux()
    serve_mux.Handle("/app/*", conf.MiddlewareMetricsIncrementor(http.StripPrefix("/app", http.FileServer(http.Dir(filepath_root)))))
    serve_mux.HandleFunc("GET /api/healthz", endpoints.ReadinessEndpointHandler)
    serve_mux.HandleFunc("GET /admin/metrics", conf.MetricsEndpointHandler)
    serve_mux.HandleFunc("POST /api/chirps", apiprocessing.ValidateChirp)
    serve_mux.HandleFunc("/api/reset", conf.MiddlewareMetricsReset)
    srv := &http.Server{
        Addr: ":" + port,
        Handler: serve_mux,
    }
    log.Printf("Serving files from %s on port: %s\n", filepath_root, port)
    log.Fatal(srv.ListenAndServe())
}
