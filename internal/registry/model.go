package registry

import (
	"sync"
	"time"
)

type Service struct {
	Name      string `json:"name"`
	URL       string `json:"url"`
	HealthURL string `json:"health_url"`
}

type RegistryManager struct {
	mu       sync.Mutex
	services map[string]*ServiceInstance
}

type ServiceInstance struct {
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	HealthURL string    `json:"health_url"`
	LastBeat  time.Time `json:"last_beat"`
	Status    string    `json:"status"`
}