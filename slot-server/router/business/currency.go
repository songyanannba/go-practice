package business

import (
	"github.com/gin-gonic/gin"
	"slot-server/api/v1"
	"slot-server/middleware"
)

type CurrencyRouter struct {
}

// InitCurrencyRouter 初始化 Currency 路由信息
func (s *CurrencyRouter) InitCurrencyRouter(Router *gin.RouterGroup) {
	currencyRouter := Router.Group("currency").Use(middleware.OperationRecord())
	currencyRouterWithoutRecord := Router.Group("currency")
	var currencyApi = v1.ApiGroupApp.BusinessApiGroup.CurrencyApi
	{
		currencyRouter.POST("createCurrency", currencyApi.CreateCurrency)             // 新建Currency
		currencyRouter.DELETE("deleteCurrency", currencyApi.DeleteCurrency)           // 删除Currency
		currencyRouter.DELETE("deleteCurrencyByIds", currencyApi.DeleteCurrencyByIds) // 批量删除Currency
		currencyRouter.PUT("updateCurrency", currencyApi.UpdateCurrency)              // 更新Currency
	}
	{
		currencyRouterWithoutRecord.GET("findCurrency", currencyApi.FindCurrency)       // 根据ID获取Currency
		currencyRouterWithoutRecord.GET("getCurrencyList", currencyApi.GetCurrencyList) // 获取Currency列表
	}
}
