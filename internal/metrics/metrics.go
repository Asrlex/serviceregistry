package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"

	"sync"
)

var (
	ServicesRegisteredTotal = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: string(TotalServices),
			Help: "Total number of services registered",
		},
	)

	ServicesHealthy = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: string(HealthyServices),
			Help: "Current number of healthy services",
		},
	)

	ServicesUnhealthy = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: string(UnhealthyServices),
			Help: "Total number of unhealthy services",
		},
	)
)

func Init() {
	sync.OnceFunc(func() {
		prometheus.MustRegister(
			ServicesRegisteredTotal,
			ServicesHealthy,
			ServicesUnhealthy,
			Uptime,
			collectors.NewGoCollector(),
			collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		)
	})
}
