package middleware

import (
	"fmt"
	"net/http"
    "log"
)

type ApiConfig struct {
    FileserverHits int
}

func ReadinessEndpointHandler(w http.ResponseWriter, d *http.Request) {
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
}

func (cfg *ApiConfig) MetricsEndpointHandler(w http.ResponseWriter, d *http.Request) {
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    output := fmt.Sprintf("Hits: %d\n", cfg.FileserverHits)
    w.Write([]byte(output))
}

func (cfg *ApiConfig) MiddlewareMetricsIncrementor(next http.Handler) http.Handler{
    return http.HandlerFunc(func(w http.ResponseWriter, d *http.Request) {
        cfg.FileserverHits++
        log.Println(cfg.FileserverHits)
        next.ServeHTTP(w, d)
    })
}

func (cfg *ApiConfig) MiddlewareMetricsReset(w http.ResponseWriter, d *http.Request) {
    cfg.FileserverHits = 0
    log.Printf("Reseting hit counter: %d\n", cfg.FileserverHits)
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.WriteHeader(http.StatusOK)
}
