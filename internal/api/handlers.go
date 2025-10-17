package api

import (
	"net/http"
	"encoding/json"

	"serviceregistry/internal/registry"
)

// healthCheckHandler responds with a simple status message to indicate the service is running
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"service": "serviceregistry",
	})
}

// registerHandler handles service registration requests
func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var s registry.Service
	var si registry.ServiceInstance
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}
	si = registry.ServiceInstance{
		Name: s.Name,
		URL:  s.URL,
	}
	rm := registry.GetRegistryManager()
	rm.RegisterService(si)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(si)
}

// deregisterHandler handles service deregistration requests
func deregisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var s registry.Service
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}
	rm := registry.GetRegistryManager()
	err = rm.DeregisterService(s.Name, s.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "service deregistered successfully",
	})
}

// serviceHandler handles requests to retrieve registered services
func serviceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Path[len("/services/"):]
	rm := registry.GetRegistryManager()
	w.Header().Set("Content-Type", "application/json")
	if name == "" {
		services, err := rm.GetAllServices()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			json.NewEncoder(w).Encode(services)
			return
		}
	}
	service, err := rm.GetService(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(service)
}