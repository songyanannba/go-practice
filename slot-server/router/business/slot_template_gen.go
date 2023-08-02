package business

import (
	"github.com/gin-gonic/gin"
	"slot-server/api/v1"
	"slot-server/middleware"
)

type SlotTemplateGenRouter struct {
}

// InitSlotTemplateGenRouter 初始化 SlotTemplateGen 路由信息
func (s *SlotTemplateGenRouter) InitSlotTemplateGenRouter(Router *gin.RouterGroup) {
	slotTemplateGenRouter := Router.Group("slotTemplateGen").Use(middleware.OperationRecord())
	slotTemplateGenRouterWithoutRecord := Router.Group("slotTemplateGen")
	var slotTemplateGenApi = v1.ApiGroupApp.BusinessApiGroup.SlotTemplateGenApi
	{
		slotTemplateGenRouter.POST("createSlotTemplateGen", slotTemplateGenApi.CreateSlotTemplateGen)             // 新建SlotTemplateGen
		slotTemplateGenRouter.DELETE("deleteSlotTemplateGen", slotTemplateGenApi.DeleteSlotTemplateGen)           // 删除SlotTemplateGen
		slotTemplateGenRouter.DELETE("deleteSlotTemplateGenByIds", slotTemplateGenApi.DeleteSlotTemplateGenByIds) // 批量删除SlotTemplateGen
		slotTemplateGenRouter.PUT("updateSlotTemplateGen", slotTemplateGenApi.UpdateSlotTemplateGen)              // 更新SlotTemplateGen
		slotTemplateGenRouter.PUT("generateSlotTemplateGen", slotTemplateGenApi.GenerateSlotTemplateGen)          // 更新SlotTemplateGen
	}
	{
		slotTemplateGenRouterWithoutRecord.GET("findSlotTemplateGen", slotTemplateGenApi.FindSlotTemplateGen)       // 根据ID获取SlotTemplateGen
		slotTemplateGenRouterWithoutRecord.GET("getSlotTemplateGenList", slotTemplateGenApi.GetSlotTemplateGenList) // 获取SlotTemplateGen列表
	}
}
