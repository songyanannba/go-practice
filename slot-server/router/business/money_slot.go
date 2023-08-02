package business

import (
	"github.com/gin-gonic/gin"
	"slot-server/api/v1"
	"slot-server/middleware"
)

type MoneySlotRouter struct {
}

// InitMoneySlotRouter 初始化 MoneySlot 路由信息
func (s *MoneySlotRouter) InitMoneySlotRouter(Router *gin.RouterGroup) {
	moneySlotRouter := Router.Group("moneySlot").Use(middleware.OperationRecord())
	moneySlotRouterWithoutRecord := Router.Group("moneySlot")
	var moneySlotApi = v1.ApiGroupApp.BusinessApiGroup.MoneySlotApi
	{
		moneySlotRouter.POST("createMoneySlot", moneySlotApi.CreateMoneySlot)             // 新建MoneySlot
		moneySlotRouter.DELETE("deleteMoneySlot", moneySlotApi.DeleteMoneySlot)           // 删除MoneySlot
		moneySlotRouter.DELETE("deleteMoneySlotByIds", moneySlotApi.DeleteMoneySlotByIds) // 批量删除MoneySlot
		moneySlotRouter.PUT("updateMoneySlot", moneySlotApi.UpdateMoneySlot)              // 更新MoneySlot
	}
	{
		moneySlotRouterWithoutRecord.GET("findMoneySlot", moneySlotApi.FindMoneySlot)       // 根据ID获取MoneySlot
		moneySlotRouterWithoutRecord.GET("getMoneySlotList", moneySlotApi.GetMoneySlotList) // 获取MoneySlot列表
		moneySlotRouterWithoutRecord.GET("recalculate", moneySlotApi.Recalculate)           // 重新计算
	}
}
