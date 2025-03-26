package api

import (
	"log"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received health check request - %v", http.StatusOK)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Sidecar injector is healthy!"))
}
