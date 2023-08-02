package system

import (
	"github.com/gin-gonic/gin"
	v1 "slot-server/api/v1"
	"slot-server/middleware"
)

type SysRouter struct{}

func (s *SysRouter) InitSystemRouter(Router *gin.RouterGroup) {
	sysRouter := Router.Group("system").Use(middleware.OperationRecord())
	systemApi := v1.ApiGroupApp.SystemApiGroup.SystemApi
	{
		sysRouter.POST("getSystemConfig", systemApi.GetSystemConfig) // 获取配置文件内容
		sysRouter.POST("setSystemConfig", systemApi.SetSystemConfig) // 设置配置文件内容
		sysRouter.POST("reloadSystem", systemApi.ReloadSystem)       // 重启服务
	}

	Router.Group("system").POST("getServerInfo", systemApi.GetServerInfo) // 获取服务器信息

}
