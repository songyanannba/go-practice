package business

import (
	"github.com/gin-gonic/gin"
	"slot-server/api/v1"
	"slot-server/middleware"
)

type MoneyLogRouter struct {
}

// InitMoneyLogRouter 初始化 MoneyLog 路由信息
func (s *MoneyLogRouter) InitMoneyLogRouter(Router *gin.RouterGroup) {
	moneyLogRouter := Router.Group("moneyLog").Use(middleware.OperationRecord())
	moneyLogRouterWithoutRecord := Router.Group("moneyLog")
	var moneyLogApi = v1.ApiGroupApp.BusinessApiGroup.MoneyLogApi
	{
		moneyLogRouter.POST("createMoneyLog", moneyLogApi.CreateMoneyLog)             // 新建MoneyLog
		moneyLogRouter.DELETE("deleteMoneyLog", moneyLogApi.DeleteMoneyLog)           // 删除MoneyLog
		moneyLogRouter.DELETE("deleteMoneyLogByIds", moneyLogApi.DeleteMoneyLogByIds) // 批量删除MoneyLog
		moneyLogRouter.PUT("updateMoneyLog", moneyLogApi.UpdateMoneyLog)              // 更新MoneyLog
	}
	{
		moneyLogRouterWithoutRecord.GET("findMoneyLog", moneyLogApi.FindMoneyLog)       // 根据ID获取MoneyLog
		moneyLogRouterWithoutRecord.GET("getMoneyLogList", moneyLogApi.GetMoneyLogList) // 获取MoneyLog列表
	}
}
