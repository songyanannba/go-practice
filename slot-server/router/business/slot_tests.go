package business

import (
	"github.com/gin-gonic/gin"
	"slot-server/api/v1"
	"slot-server/middleware"
)

type SlotTestsRouter struct {
}

// InitSlotTestsRouter 初始化 SlotTests 路由信息
func (s *SlotTestsRouter) InitSlotTestsRouter(Router *gin.RouterGroup) {
	slotTestsRouter := Router.Group("slotTests").Use(middleware.OperationRecord())
	slotTestsRouterWithoutRecord := Router.Group("slotTests")
	var slotTestsApi = v1.ApiGroupApp.BusinessApiGroup.SlotTestsApi
	{
		slotTestsRouter.POST("createSlotTests", slotTestsApi.CreateSlotTests)             // 新建SlotTests
		slotTestsRouter.DELETE("deleteSlotTests", slotTestsApi.DeleteSlotTests)           // 删除SlotTests
		slotTestsRouter.DELETE("deleteSlotTestsByIds", slotTestsApi.DeleteSlotTestsByIds) // 批量删除SlotTests
		slotTestsRouter.PUT("updateSlotTests", slotTestsApi.UpdateSlotTests)              // 更新SlotTests
		slotTestsRouter.DELETE("truncateSlotTests", slotTestsApi.Truncate)                // 根据ID获取SlotTests
	}
	{
		slotTestsRouterWithoutRecord.GET("findSlotTests", slotTestsApi.FindSlotTests)       // 根据ID获取SlotTests
		slotTestsRouterWithoutRecord.GET("getSlotTestsList", slotTestsApi.GetSlotTestsList) // 获取SlotTests列表
	}
}
