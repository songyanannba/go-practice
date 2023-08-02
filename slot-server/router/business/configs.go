package business

import (
	"github.com/gin-gonic/gin"
	"slot-server/api/v1"
	"slot-server/middleware"
)

type ConfigsRouter struct {
}

// InitConfigsRouter 初始化 Configs 路由信息
func (s *ConfigsRouter) InitConfigsRouter(Router *gin.RouterGroup) {
	configsRouter := Router.Group("configs").Use(middleware.OperationRecord())
	configsRouterWithoutRecord := Router.Group("configs")
	var configsApi = v1.ApiGroupApp.BusinessApiGroup.ConfigsApi
	{
		configsRouter.POST("createConfigs", configsApi.CreateConfigs)             // 新建Configs
		configsRouter.DELETE("deleteConfigs", configsApi.DeleteConfigs)           // 删除Configs
		configsRouter.DELETE("deleteConfigsByIds", configsApi.DeleteConfigsByIds) // 批量删除Configs
		configsRouter.PUT("updateConfigs", configsApi.UpdateConfigs)              // 更新Configs
	}
	{
		configsRouterWithoutRecord.GET("findConfigs", configsApi.FindConfigs)       // 根据ID获取Configs
		configsRouterWithoutRecord.GET("getConfigsList", configsApi.GetConfigsList) // 获取Configs列表
	}
}
