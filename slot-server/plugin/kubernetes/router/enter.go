package router

type RouterGroup struct {
	ClusterRouter
	ProxyApiRouter
	MetricRouter
	WsApiRouter
}

var RouterGroupApp = new(RouterGroup)
