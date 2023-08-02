package router

import (
	"github.com/gin-gonic/gin"
	"slot-server/middleware"
	"slot-server/plugin/kubernetes/api"
)

type MetricRouter struct{}

func (u *MetricRouter) InitMetricRouter(Router *gin.RouterGroup) {
	metricRouter := Router.Group("metrics").Use(middleware.OperationRecord())
	metricsApi := api.ApiGroupApp.MetricsApi
	{
		metricRouter.POST("get", metricsApi.MetricsGet) // 监控数据获取
	}
}
