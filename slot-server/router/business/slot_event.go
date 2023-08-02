package business

import (
	"slot-server/api/v1"
	"slot-server/middleware"
	"github.com/gin-gonic/gin"
)

type SlotEventRouter struct {
}

// InitSlotEventRouter 初始化 SlotEvent 路由信息
func (s *SlotEventRouter) InitSlotEventRouter(Router *gin.RouterGroup) {
	slotEventRouter := Router.Group("slotEvent").Use(middleware.OperationRecord())
	slotEventRouterWithoutRecord := Router.Group("slotEvent")
	var slotEventApi = v1.ApiGroupApp.BusinessApiGroup.SlotEventApi
	{
		slotEventRouter.POST("createSlotEvent", slotEventApi.CreateSlotEvent)   // 新建SlotEvent
		slotEventRouter.DELETE("deleteSlotEvent", slotEventApi.DeleteSlotEvent) // 删除SlotEvent
		slotEventRouter.DELETE("deleteSlotEventByIds", slotEventApi.DeleteSlotEventByIds) // 批量删除SlotEvent
		slotEventRouter.PUT("updateSlotEvent", slotEventApi.UpdateSlotEvent)    // 更新SlotEvent
	}
	{
		slotEventRouterWithoutRecord.GET("findSlotEvent", slotEventApi.FindSlotEvent)        // 根据ID获取SlotEvent
		slotEventRouterWithoutRecord.GET("getSlotEventList", slotEventApi.GetSlotEventList)  // 获取SlotEvent列表
	}
}
