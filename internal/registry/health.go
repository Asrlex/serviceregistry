package registry

import (
	"log"
	"net/http"
	"time"

	"serviceregistry/internal/metrics"
)

func (rm *RegistryManager) StartHealthCheck(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			rm.CheckAllServices()
		}
	}()
}

func (rm *RegistryManager) CheckAllServices() {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	metrics.ServicesUnhealthy.Set(0)
	metrics.ServicesHealthy.Set(0)
	for name, instance := range rm.services {
		resp, callErr := http.Get(instance.HealthURL)
		if callErr != nil || resp.StatusCode != http.StatusOK {
			log.Printf("Service %s at %s is unhealthy", name, instance.URL)
			instance.Status = "unhealthy"
			metrics.ServicesHealthy.Dec()
			metrics.ServicesUnhealthy.Inc()
		} else {
			log.Printf("Service %s at %s is healthy", name, instance.URL)
			instance.Status = "healthy"
			instance.LastBeat = time.Now()
			metrics.ServicesHealthy.Inc()
			metrics.ServicesUnhealthy.Dec()
		}
	}
}