package registry

import (
	"errors"
	"log"
	"time"

	"serviceregistry/internal/metrics"
)

var RegistryManagerInstance *RegistryManager

// GetRegistryManager returns the singleton instance of RegistryManager.
func GetRegistryManager() *RegistryManager {
	if RegistryManagerInstance != nil {
		return RegistryManagerInstance
	}
	RegistryManagerInstance = NewRegistryManager()
	return RegistryManagerInstance
}

// NewRegistryManager creates a new instance of RegistryManager.
func NewRegistryManager() *RegistryManager {
	return &RegistryManager{
		services: make(map[string]*ServiceInstance),
	}
}

// RegisterService adds a new service instance to the registry.
func (rm *RegistryManager) RegisterService(instance ServiceInstance) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	instance.LastBeat = time.Now()
	instance.Status = "healthy"
	rm.services[instance.Name] = &instance
	metrics.ServicesRegisteredTotal.Inc()
	metrics.ServicesHealthy.Inc()
	log.Printf("Registered service instance: %s at %s", instance.Name, instance.URL)
}

// DeregisterService removes a service instance from the registry.
func (rm *RegistryManager) DeregisterService(name, url string) error {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	if _, exists := rm.services[name]; !exists {
		log.Printf("No instances found for service: %s", name)
		return errors.New("service not found")
	}

	delete(rm.services, name)
	metrics.ServicesRegisteredTotal.Dec()
	metrics.ServicesHealthy.Dec()
	log.Printf("Deregistered service instance: %s at %s", name, url)
	return nil
}

// GetAllServices retrieves all registered service instances.
func (rm *RegistryManager) GetAllServices() ([]ServiceInstance, error) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	services := make([]ServiceInstance, 0, len(rm.services))
	for _, instance := range rm.services {
		services = append(services, *instance)
	}
	return services, nil
}

// GetService retrieves a service instance by name.
func (rm *RegistryManager) GetService(name string) (ServiceInstance, error) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	
	instance, exists := rm.services[name]
	if !exists {
		return ServiceInstance{}, errors.New("service not found")
	}
	return *instance, nil
}