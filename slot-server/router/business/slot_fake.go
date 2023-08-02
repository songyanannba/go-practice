package business

import (
	"github.com/gin-gonic/gin"
	"slot-server/api/v1"
	"slot-server/middleware"
)

type SlotFakeRouter struct {
}

// InitSlotFakeRouter 初始化 SlotFake 路由信息
func (s *SlotFakeRouter) InitSlotFakeRouter(Router *gin.RouterGroup) {
	slotFakeRouter := Router.Group("slotFake").Use(middleware.OperationRecord())
	slotFakeRouterWithoutRecord := Router.Group("slotFake")
	var slotFakeApi = v1.ApiGroupApp.BusinessApiGroup.SlotFakeApi
	{
		slotFakeRouter.POST("createSlotFake", slotFakeApi.CreateSlotFake)             // 新建SlotFake
		slotFakeRouter.DELETE("deleteSlotFake", slotFakeApi.DeleteSlotFake)           // 删除SlotFake
		slotFakeRouter.DELETE("deleteSlotFakeByIds", slotFakeApi.DeleteSlotFakeByIds) // 批量删除SlotFake
		slotFakeRouter.PUT("updateSlotFake", slotFakeApi.UpdateSlotFake)              // 更新SlotFake
	}
	{
		slotFakeRouterWithoutRecord.GET("findSlotFake", slotFakeApi.FindSlotFake)       // 根据ID获取SlotFake
		slotFakeRouterWithoutRecord.GET("getSlotFakeList", slotFakeApi.GetSlotFakeList) // 获取SlotFake列表
	}
}
