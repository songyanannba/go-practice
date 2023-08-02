package business

import (
	"github.com/gin-gonic/gin"
	"slot-server/api/v1"
	"slot-server/middleware"
)

type SlotUserSpinRouter struct {
}

// InitSlotUserSpinRouter 初始化 SlotUserSpin 路由信息
func (s *SlotUserSpinRouter) InitSlotUserSpinRouter(Router *gin.RouterGroup) {
	slotUserSpinRouter := Router.Group("slotUserSpin").Use(middleware.OperationRecord())
	slotUserSpinRouterWithoutRecord := Router.Group("slotUserSpin")
	var slotUserSpinApi = v1.ApiGroupApp.BusinessApiGroup.SlotUserSpinApi
	{
		slotUserSpinRouter.POST("createSlotUserSpin", slotUserSpinApi.CreateSlotUserSpin)             // 新建SlotUserSpin
		slotUserSpinRouter.DELETE("deleteSlotUserSpin", slotUserSpinApi.DeleteSlotUserSpin)           // 删除SlotUserSpin
		slotUserSpinRouter.DELETE("deleteSlotUserSpinByIds", slotUserSpinApi.DeleteSlotUserSpinByIds) // 批量删除SlotUserSpin
		slotUserSpinRouter.PUT("updateSlotUserSpin", slotUserSpinApi.UpdateSlotUserSpin)              // 更新SlotUserSpin
	}
	{
		slotUserSpinRouterWithoutRecord.GET("findSlotUserSpin", slotUserSpinApi.FindSlotUserSpin)       // 根据ID获取SlotUserSpin
		slotUserSpinRouterWithoutRecord.GET("getSlotUserSpinList", slotUserSpinApi.GetSlotUserSpinList) // 获取SlotUserSpin列表
	}
}
