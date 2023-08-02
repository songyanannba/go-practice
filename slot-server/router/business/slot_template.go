package business

import (
	"github.com/gin-gonic/gin"
	"slot-server/api/v1"
	"slot-server/middleware"
)

type SlotTemplateRouter struct {
}

// InitSlotTemplateRouter 初始化 SlotTemplate 路由信息
func (s *SlotTemplateRouter) InitSlotTemplateRouter(Router *gin.RouterGroup) {
	slotTemplateRouter := Router.Group("slotTemplate").Use(middleware.OperationRecord())
	slotTemplateRouterWithoutRecord := Router.Group("slotTemplate")
	var slotTemplateApi = v1.ApiGroupApp.BusinessApiGroup.SlotTemplateApi
	{
		slotTemplateRouter.POST("createSlotTemplate", slotTemplateApi.CreateSlotTemplate)             // 新建SlotTemplate
		slotTemplateRouter.DELETE("deleteSlotTemplate", slotTemplateApi.DeleteSlotTemplate)           // 删除SlotTemplate
		slotTemplateRouter.DELETE("deleteSlotTemplateByIds", slotTemplateApi.DeleteSlotTemplateByIds) // 批量删除SlotTemplate
		slotTemplateRouter.PUT("updateSlotTemplate", slotTemplateApi.UpdateSlotTemplate)              // 更新SlotTemplate
	}
	{
		slotTemplateRouterWithoutRecord.GET("findSlotTemplate", slotTemplateApi.FindSlotTemplate)       // 根据ID获取SlotTemplate
		slotTemplateRouterWithoutRecord.GET("getSlotTemplateList", slotTemplateApi.GetSlotTemplateList) // 获取SlotTemplate列表
	}
}
