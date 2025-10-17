package main

import (
	"github.com/joho/godotenv"
	
	"log"
	"net/http"
	"time"

	"serviceregistry/internal/api"
	"serviceregistry/internal/registry"
)

func main() {
	godotenv.Load()
	rm := registry.GetRegistryManager()
	rm.StartHealthCheck(30 * time.Second)
	router := api.NewRouter()
	
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}