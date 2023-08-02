package router

import (
	"github.com/gin-gonic/gin"
	"slot-server/plugin/kubernetes/api"
)

type WsApiRouter struct{}

func (t *WsApiRouter) InitWsRouter(Router *gin.RouterGroup) {
	WsRouter := Router.Group("kubernetes/pods")
	wsApi := api.ApiGroupApp.WsApi
	WsRouter.GET("/terminal", wsApi.Terminal) // 终端
	WsRouter.GET("/logs", wsApi.ContainerLog) // 终端日志
}
