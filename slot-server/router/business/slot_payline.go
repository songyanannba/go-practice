package business

import (
	"github.com/gin-gonic/gin"
	"slot-server/api/v1"
	"slot-server/middleware"
)

type SlotPaylineRouter struct {
}

// InitSlotPaylineRouter 初始化 SlotPayline 路由信息
func (s *SlotPaylineRouter) InitSlotPaylineRouter(Router *gin.RouterGroup) {
	slotPaylineRouter := Router.Group("slotPayline").Use(middleware.OperationRecord())
	slotPaylineRouterWithoutRecord := Router.Group("slotPayline")
	var slotPaylineApi = v1.ApiGroupApp.BusinessApiGroup.SlotPaylineApi
	{
		slotPaylineRouter.POST("createSlotPayline", slotPaylineApi.CreateSlotPayline)             // 新建SlotPayline
		slotPaylineRouter.DELETE("deleteSlotPayline", slotPaylineApi.DeleteSlotPayline)           // 删除SlotPayline
		slotPaylineRouter.DELETE("deleteSlotPaylineByIds", slotPaylineApi.DeleteSlotPaylineByIds) // 批量删除SlotPayline
		slotPaylineRouter.PUT("updateSlotPayline", slotPaylineApi.UpdateSlotPayline)              // 更新SlotPayline
	}
	{
		slotPaylineRouterWithoutRecord.GET("findSlotPayline", slotPaylineApi.FindSlotPayline)       // 根据ID获取SlotPayline
		slotPaylineRouterWithoutRecord.GET("getSlotPaylineList", slotPaylineApi.GetSlotPaylineList) // 获取SlotPayline列表
	}
}
