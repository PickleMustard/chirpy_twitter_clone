package main

import (
    "log"
	"net/http"
)

func main() {
    const filepath_root = "."
    const port = "8080"

    serve_mux := http.NewServeMux()
    serve_mux.Handle("/", http.FileServer(http.Dir(filepath_root)))
    srv := &http.Server{
        Addr: ":" + port,
        Handler: serve_mux,
    }
    log.Printf("Serving files from %s on port: %s\n", filepath_root, port)
    log.Fatal(srv.ListenAndServe())
}
