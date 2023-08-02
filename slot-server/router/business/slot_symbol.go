package business

import (
	"github.com/gin-gonic/gin"
	"slot-server/api/v1"
	"slot-server/middleware"
)

type SlotSymbolRouter struct {
}

// InitSlotSymbolRouter 初始化 SlotSymbol 路由信息
func (s *SlotSymbolRouter) InitSlotSymbolRouter(Router *gin.RouterGroup) {
	slotSymbolRouter := Router.Group("slotSymbol").Use(middleware.OperationRecord())
	slotSymbolRouterWithoutRecord := Router.Group("slotSymbol")
	var slotSymbolApi = v1.ApiGroupApp.BusinessApiGroup.SlotSymbolApi
	{
		slotSymbolRouter.POST("createSlotSymbol", slotSymbolApi.CreateSlotSymbol)             // 新建SlotSymbol
		slotSymbolRouter.DELETE("deleteSlotSymbol", slotSymbolApi.DeleteSlotSymbol)           // 删除SlotSymbol
		slotSymbolRouter.DELETE("deleteSlotSymbolByIds", slotSymbolApi.DeleteSlotSymbolByIds) // 批量删除SlotSymbol
		slotSymbolRouter.PUT("updateSlotSymbol", slotSymbolApi.UpdateSlotSymbol)              // 更新SlotSymbol
	}
	{
		slotSymbolRouterWithoutRecord.GET("findSlotSymbol", slotSymbolApi.FindSlotSymbol)       // 根据ID获取SlotSymbol
		slotSymbolRouterWithoutRecord.GET("getSlotSymbolList", slotSymbolApi.GetSlotSymbolList) // 获取SlotSymbol列表
	}
}
