package business

import (
	"github.com/gin-gonic/gin"
	"slot-server/api/v1"
	"slot-server/middleware"
)

type DebugConfigRouter struct {
}

// InitDebugConfigRouter 初始化 DebugConfig 路由信息
func (s *DebugConfigRouter) InitDebugConfigRouter(Router *gin.RouterGroup) {
	debugConfigRouter := Router.Group("debugConfig").Use(middleware.OperationRecord())
	debugConfigRouterWithoutRecord := Router.Group("debugConfig")
	var debugConfigApi = v1.ApiGroupApp.BusinessApiGroup.DebugConfigApi
	{
		debugConfigRouter.POST("createDebugConfig", debugConfigApi.CreateDebugConfig)             // 新建DebugConfig
		debugConfigRouter.DELETE("deleteDebugConfig", debugConfigApi.DeleteDebugConfig)           // 删除DebugConfig
		debugConfigRouter.DELETE("deleteDebugConfigByIds", debugConfigApi.DeleteDebugConfigByIds) // 批量删除DebugConfig
		debugConfigRouter.PUT("updateDebugConfig", debugConfigApi.UpdateDebugConfig)              // 更新DebugConfig
	}
	{
		debugConfigRouterWithoutRecord.GET("findDebugConfig", debugConfigApi.FindDebugConfig)       // 根据ID获取DebugConfig
		debugConfigRouterWithoutRecord.GET("getDebugConfigList", debugConfigApi.GetDebugConfigList) // 获取DebugConfig列表
		debugConfigRouterWithoutRecord.GET("getSlotTags", debugConfigApi.GetSlotTags)               // 获取DebugConfig列表
	}
}
