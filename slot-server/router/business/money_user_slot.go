package business

import (
	"github.com/gin-gonic/gin"
	"slot-server/api/v1"
	"slot-server/middleware"
)

type MoneyUserSlotRouter struct {
}

// InitMoneyUserSlotRouter 初始化 MoneyUserSlot 路由信息
func (s *MoneyUserSlotRouter) InitMoneyUserSlotRouter(Router *gin.RouterGroup) {
	moneyUserSlotRouter := Router.Group("moneyUserSlot").Use(middleware.OperationRecord())
	moneyUserSlotRouterWithoutRecord := Router.Group("moneyUserSlot")
	var moneyUserSlotApi = v1.ApiGroupApp.BusinessApiGroup.MoneyUserSlotApi
	{
		moneyUserSlotRouter.POST("createMoneyUserSlot", moneyUserSlotApi.CreateMoneyUserSlot)             // 新建MoneyUserSlot
		moneyUserSlotRouter.DELETE("deleteMoneyUserSlot", moneyUserSlotApi.DeleteMoneyUserSlot)           // 删除MoneyUserSlot
		moneyUserSlotRouter.DELETE("deleteMoneyUserSlotByIds", moneyUserSlotApi.DeleteMoneyUserSlotByIds) // 批量删除MoneyUserSlot
		moneyUserSlotRouter.PUT("updateMoneyUserSlot", moneyUserSlotApi.UpdateMoneyUserSlot)              // 更新MoneyUserSlot
	}
	{
		moneyUserSlotRouterWithoutRecord.GET("findMoneyUserSlot", moneyUserSlotApi.FindMoneyUserSlot)       // 根据ID获取MoneyUserSlot
		moneyUserSlotRouterWithoutRecord.GET("getMoneyUserSlotList", moneyUserSlotApi.GetMoneyUserSlotList) // 获取MoneyUserSlot列表
		moneyUserSlotRouterWithoutRecord.GET("recalculate", moneyUserSlotApi.Recalculate)
	}
}
