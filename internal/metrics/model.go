package metrics

type MetricName string

const (
	TotalServices         MetricName = "services_registered_total"
	HealthyServices        MetricName = "services_healthy"
	UnhealthyServices      MetricName = "services_unhealthy"
)
