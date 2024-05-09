package middleware

import (
	//"log"
	"net/http"
    //"fmt"
)

type api_config struct {
    fileserver_hits int
}

func readiness_endpoint_handler(w http.ResponseWriter, d *http.Request) {
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.WriteHeader(http.StatusOK)

    //body_text := fmt.Sprintf("Status: %d", http.StatusOK)

    w.Write([]byte("OK"))
}

func (cfg *api_config) middleware_metrics_incrementor_generator() (f func(next http.Handler) http.Handler) {
    var conf api_config;
    conf.fileserver_hits = 0;

    increment_handler := func(next http.Handler) http.Handler {
        conf.fileserver_hits++
        return readiness_endpoint_handler()
    }

    return increment_handler
}
