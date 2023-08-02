package business

import (
	"github.com/gin-gonic/gin"
	"slot-server/api/v1"
	"slot-server/middleware"
)

type SlotReelDataRouter struct {
}

// InitSlotReelDataRouter 初始化 SlotReelData 路由信息
func (s *SlotReelDataRouter) InitSlotReelDataRouter(Router *gin.RouterGroup) {
	slotReelDataRouter := Router.Group("slotReelData").Use(middleware.OperationRecord())
	slotReelDataRouterWithoutRecord := Router.Group("slotReelData")
	var slotReelDataApi = v1.ApiGroupApp.BusinessApiGroup.SlotReelDataApi
	{
		slotReelDataRouter.POST("createSlotReelData", slotReelDataApi.CreateSlotReelData)             // 新建SlotReelData
		slotReelDataRouter.DELETE("deleteSlotReelData", slotReelDataApi.DeleteSlotReelData)           // 删除SlotReelData
		slotReelDataRouter.DELETE("deleteSlotReelDataByIds", slotReelDataApi.DeleteSlotReelDataByIds) // 批量删除SlotReelData
		slotReelDataRouter.PUT("updateSlotReelData", slotReelDataApi.UpdateSlotReelData)              // 更新SlotReelData
	}
	{
		slotReelDataRouterWithoutRecord.GET("findSlotReelData", slotReelDataApi.FindSlotReelData)       // 根据ID获取SlotReelData
		slotReelDataRouterWithoutRecord.GET("getSlotReelDataList", slotReelDataApi.GetSlotReelDataList) // 获取SlotReelData列表
	}
}
