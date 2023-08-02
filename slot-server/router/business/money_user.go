package business

import (
	"github.com/gin-gonic/gin"
	"slot-server/api/v1"
	"slot-server/middleware"
)

type MoneyUserRouter struct {
}

// InitMoneyUserRouter 初始化 MoneyUser 路由信息
func (s *MoneyUserRouter) InitMoneyUserRouter(Router *gin.RouterGroup) {
	MoneyUserRouter := Router.Group("moneyUser").Use(middleware.OperationRecord())
	moneyUserRouterWithoutRecord := Router.Group("moneyUser")
	var moneyUserApi = v1.ApiGroupApp.BusinessApiGroup.MoneyUserApi
	{
		MoneyUserRouter.POST("createMoneyUser", moneyUserApi.CreateMoneyUser)             // 新建MoneyUser
		MoneyUserRouter.DELETE("deleteMoneyUser", moneyUserApi.DeleteMoneyUser)           // 删除MoneyUser
		MoneyUserRouter.DELETE("deleteMoneyUserByIds", moneyUserApi.DeleteMoneyUserByIds) // 批量删除MoneyUser
		MoneyUserRouter.PUT("updateMoneyUser", moneyUserApi.UpdateMoneyUser)              // 更新MoneyUser
	}
	{
		moneyUserRouterWithoutRecord.GET("findMoneyUser", moneyUserApi.FindMoneyUser)       // 根据ID获取MoneyUser
		moneyUserRouterWithoutRecord.GET("getMoneyUserList", moneyUserApi.GetMoneyUserList) // 获取MoneyUser列表
		MoneyUserRouter.GET("recalculate", moneyUserApi.Recalculate)                        // 重新计算
	}
}
