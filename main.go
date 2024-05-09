package main

import (
	"log"
	"net/http"
    "internal/middleware"
    //"fmt"
)

//Handler Function for a readiness endpoint
//Matches the function signature of http.HandlerFunc

func main() {
    const filepath_root = "."
    const port = "8080"

    middleware_metrics_incrementor := middleware.middleware_metrics_incrementor_generator()

    serve_mux := http.NewServeMux()
    serve_mux.Handle("/app/*", http.StripPrefix("/app", http.FileServer(http.Dir(filepath_root))))
    serve_mux.HandleFunc("/healthz", middleware.middleware_metrics_incrementor)
    srv := &http.Server{
        Addr: ":" + port,
        Handler: serve_mux,
    }
    log.Printf("Serving files from %s on port: %s\n", filepath_root, port)
    log.Fatal(srv.ListenAndServe())
}
