package business

import (
	"github.com/gin-gonic/gin"
	"slot-server/api/v1"
	"slot-server/middleware"
)

type MerchantRouter struct {
}

// InitMerchantRouter 初始化 Merchant 路由信息
func (s *MerchantRouter) InitMerchantRouter(Router *gin.RouterGroup) {
	merchantRouter := Router.Group("merchant").Use(middleware.OperationRecord())
	merchantRouterWithoutRecord := Router.Group("merchant")
	var merchantApi = v1.ApiGroupApp.BusinessApiGroup.MerchantApi
	{
		merchantRouter.POST("createMerchant", merchantApi.CreateMerchant)             // 新建Merchant
		merchantRouter.DELETE("deleteMerchant", merchantApi.DeleteMerchant)           // 删除Merchant
		merchantRouter.DELETE("deleteMerchantByIds", merchantApi.DeleteMerchantByIds) // 批量删除Merchant
		merchantRouter.PUT("updateMerchant", merchantApi.UpdateMerchant)              // 更新Merchant
	}
	{
		merchantRouterWithoutRecord.GET("findMerchant", merchantApi.FindMerchant)       // 根据ID获取Merchant
		merchantRouterWithoutRecord.GET("getMerchantList", merchantApi.GetMerchantList) // 获取Merchant列表
	}
}
