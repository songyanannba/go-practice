package business

import (
	"github.com/gin-gonic/gin"
	"slot-server/api/v1"
	"slot-server/middleware"
)

type SlotGenTplRouter struct {
}

// InitSlotGenTplRouter 初始化 SlotGenTpl 路由信息
func (s *SlotGenTplRouter) InitSlotGenTplRouter(Router *gin.RouterGroup) {
	slotGenTplRouter := Router.Group("slotGenTpl").Use(middleware.OperationRecord())
	slotGenTplRouterWithoutRecord := Router.Group("slotGenTpl")
	var slotGenTplApi = v1.ApiGroupApp.BusinessApiGroup.SlotGenTplApi
	{
		slotGenTplRouter.POST("createSlotGenTpl", slotGenTplApi.CreateSlotGenTpl)             // 新建SlotGenTpl
		slotGenTplRouter.DELETE("deleteSlotGenTpl", slotGenTplApi.DeleteSlotGenTpl)           // 删除SlotGenTpl
		slotGenTplRouter.DELETE("deleteSlotGenTplByIds", slotGenTplApi.DeleteSlotGenTplByIds) // 批量删除SlotGenTpl
		slotGenTplRouter.PUT("updateSlotGenTpl", slotGenTplApi.UpdateSlotGenTpl)              // 更新SlotGenTpl
	}
	{
		slotGenTplRouterWithoutRecord.GET("findSlotGenTpl", slotGenTplApi.FindSlotGenTpl)       // 根据ID获取SlotGenTpl
		slotGenTplRouterWithoutRecord.GET("getSlotGenTplList", slotGenTplApi.GetSlotGenTplList) // 获取SlotGenTpl列表
	}
}
