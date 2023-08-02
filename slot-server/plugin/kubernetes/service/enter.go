package service

type ServiceGroup struct {
	ClusterService
	ProxyService
	MetricsService
}

var ServiceGroupApp = new(ServiceGroup)
