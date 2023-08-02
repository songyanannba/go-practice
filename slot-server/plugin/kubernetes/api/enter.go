package api

import "slot-server/plugin/kubernetes/service"

type ApiGroup struct {
	ClustersApi
	ProxyApi
	MetricsApi
	WsApi
}

var ApiGroupApp = new(ApiGroup)

var (
	clusterService = service.ServiceGroupApp.ClusterService
	metricService  = service.ServiceGroupApp.MetricsService
	proxyService   = service.ServiceGroupApp.ProxyService
)
