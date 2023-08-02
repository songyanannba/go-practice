package business

import (
	"github.com/gin-gonic/gin"
	"slot-server/api/v1"
	"slot-server/middleware"
)

type SlotRouter struct {
}

// InitSlotRouter 初始化 Slot 路由信息
func (s *SlotRouter) InitSlotRouter(Router *gin.RouterGroup) {
	slotRouter := Router.Group("slot").Use(middleware.OperationRecord())
	slotRouterWithoutRecord := Router.Group("slot")
	var slotApi = v1.ApiGroupApp.BusinessApiGroup.SlotApi
	{
		slotRouter.POST("createSlot", slotApi.CreateSlot)             // 新建Slot
		slotRouter.DELETE("deleteSlot", slotApi.DeleteSlot)           // 删除Slot
		slotRouter.DELETE("deleteSlotByIds", slotApi.DeleteSlotByIds) // 批量删除Slot
		slotRouter.PUT("updateSlot", slotApi.UpdateSlot)              // 更新Slot
		slotRouter.PUT("updateAllConfig", slotApi.UpdateAllConfig)    // 更新Slot所有配置
		slotRouter.POST("backendOperate", slotApi.BackendOperate)     // 请求集群操作
	}
	{
		slotRouterWithoutRecord.GET("findSlot", slotApi.FindSlot)       // 根据ID获取Slot
		slotRouterWithoutRecord.GET("getSlotList", slotApi.GetSlotList) // 获取Slot列表

	}
}
