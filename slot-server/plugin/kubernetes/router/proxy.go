package router

import (
	"github.com/gin-gonic/gin"
	"slot-server/middleware"
	"slot-server/plugin/kubernetes/api"
)

type ProxyApiRouter struct{}

func (r *ProxyApiRouter) InitProxyRouter(Router *gin.RouterGroup) {
	ProxyRouter := Router.Group("proxy").Use(middleware.OperationRecord())
	ProxyApi := api.ApiGroupApp.ProxyApi
	ProxyRouter.Any("/:cluster_id/*path", ProxyApi.K8sAPIProxy) // 资源代理
}
