package business

import (
	"github.com/gin-gonic/gin"
	"slot-server/api/v1"
	"slot-server/middleware"
)

type TrackingRouter struct {
}

// InitTrackingRouter 初始化 Tracking 路由信息
func (s *TrackingRouter) InitTrackingRouter(Router *gin.RouterGroup) {
	trackingRouter := Router.Group("tracking").Use(middleware.OperationRecord())
	trackingRouterWithoutRecord := Router.Group("tracking")
	var trackingApi = v1.ApiGroupApp.BusinessApiGroup.TrackingApi
	{
		trackingRouter.POST("createTracking", trackingApi.CreateTracking)             // 新建Tracking
		trackingRouter.DELETE("deleteTracking", trackingApi.DeleteTracking)           // 删除Tracking
		trackingRouter.DELETE("deleteTrackingByIds", trackingApi.DeleteTrackingByIds) // 批量删除Tracking
		trackingRouter.PUT("updateTracking", trackingApi.UpdateTracking)              // 更新Tracking
	}
	{
		trackingRouterWithoutRecord.GET("findTracking", trackingApi.FindTracking)       // 根据ID获取Tracking
		trackingRouterWithoutRecord.GET("getTrackingList", trackingApi.GetTrackingList) // 获取Tracking列表
	}
}
