package business

import (
	"github.com/gin-gonic/gin"
	"slot-server/api/v1"
	"slot-server/middleware"
)

type SlotPayTableRouter struct {
}

// InitSlotPayTableRouter 初始化 SlotPayTable 路由信息
func (s *SlotPayTableRouter) InitSlotPayTableRouter(Router *gin.RouterGroup) {
	slotPayTableRouter := Router.Group("slotPayTable").Use(middleware.OperationRecord())
	slotPayTableRouterWithoutRecord := Router.Group("slotPayTable")
	var slotPayTableApi = v1.ApiGroupApp.BusinessApiGroup.SlotPayTableApi
	{
		slotPayTableRouter.POST("createSlotPayTable", slotPayTableApi.CreateSlotPayTable)             // 新建SlotPayTable
		slotPayTableRouter.DELETE("deleteSlotPayTable", slotPayTableApi.DeleteSlotPayTable)           // 删除SlotPayTable
		slotPayTableRouter.DELETE("deleteSlotPayTableByIds", slotPayTableApi.DeleteSlotPayTableByIds) // 批量删除SlotPayTable
		slotPayTableRouter.PUT("updateSlotPayTable", slotPayTableApi.UpdateSlotPayTable)              // 更新SlotPayTable
	}
	{
		slotPayTableRouterWithoutRecord.GET("findSlotPayTable", slotPayTableApi.FindSlotPayTable)       // 根据ID获取SlotPayTable
		slotPayTableRouterWithoutRecord.GET("getSlotPayTableList", slotPayTableApi.GetSlotPayTableList) // 获取SlotPayTable列表
	}
}
