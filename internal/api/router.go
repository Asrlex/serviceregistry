package api

import (	
	"net/http"

	"serviceregistry/internal/auth"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthcheck", healthCheckHandler)
	mux.HandleFunc("/register", registerHandler)
	mux.HandleFunc("/deregister", deregisterHandler)
	mux.HandleFunc("/services", serviceHandler)
	mux.HandleFunc("/services/:name", serviceHandler)
	mux.Handle("/metrics", promhttp.Handler())

	return loggingMiddleware(auth.ValidateAPIKey(mux))
}
