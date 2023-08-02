package business

import (
	"slot-server/api/v1"
	"slot-server/middleware"
	"github.com/gin-gonic/gin"
)

type ApiLogRouter struct {
}

// InitApiLogRouter 初始化 ApiLog 路由信息
func (s *ApiLogRouter) InitApiLogRouter(Router *gin.RouterGroup) {
	apiLogRouter := Router.Group("apiLog").Use(middleware.OperationRecord())
	apiLogRouterWithoutRecord := Router.Group("apiLog")
	var apiLogApi = v1.ApiGroupApp.BusinessApiGroup.ApiLogApi
	{
		apiLogRouter.POST("createApiLog", apiLogApi.CreateApiLog)   // 新建ApiLog
		apiLogRouter.DELETE("deleteApiLog", apiLogApi.DeleteApiLog) // 删除ApiLog
		apiLogRouter.DELETE("deleteApiLogByIds", apiLogApi.DeleteApiLogByIds) // 批量删除ApiLog
		apiLogRouter.PUT("updateApiLog", apiLogApi.UpdateApiLog)    // 更新ApiLog
	}
	{
		apiLogRouterWithoutRecord.GET("findApiLog", apiLogApi.FindApiLog)        // 根据ID获取ApiLog
		apiLogRouterWithoutRecord.GET("getApiLogList", apiLogApi.GetApiLogList)  // 获取ApiLog列表
	}
}
