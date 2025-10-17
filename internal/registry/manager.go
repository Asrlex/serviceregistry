package registry

import (
	"errors"
	"log"
	"time"
)

var RegistryManagerInstance *RegistryManager

func GetRegistryManager() *RegistryManager {
	if RegistryManagerInstance != nil {
		return RegistryManagerInstance
	}
	RegistryManagerInstance = NewRegistryManager()
	return RegistryManagerInstance
}

func NewRegistryManager() *RegistryManager {
	return &RegistryManager{
		services: make(map[string]*ServiceInstance),
	}
}

func (rm *RegistryManager) RegisterService(instance ServiceInstance) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	instance.LastBeat = time.Now()
	instance.Status = "healthy"
	rm.services[instance.Name] = &instance
	log.Printf("Registered service instance: %s at %s", instance.Name, instance.URL)
}

func (rm *RegistryManager) DeregisterService(name, url string) error {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	if _, exists := rm.services[name]; !exists {
		log.Printf("No instances found for service: %s", name)
		return errors.New("service not found")
	}

	delete(rm.services, name)
	log.Printf("Deregistered service instance: %s at %s", name, url)
	return nil
}

func (rm *RegistryManager) GetAllServices() ([]ServiceInstance, error) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	services := make([]ServiceInstance, 0, len(rm.services))
	for _, instance := range rm.services {
		services = append(services, *instance)
	}
	return services, nil
}

func (rm *RegistryManager) GetService(name string) (ServiceInstance, error) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	
	instance, exists := rm.services[name]
	if !exists {
		return ServiceInstance{}, errors.New("service not found")
	}
	return *instance, nil
}