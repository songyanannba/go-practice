package business

import (
	"github.com/gin-gonic/gin"
	"slot-server/api/v1"
	"slot-server/middleware"
)

type JackpotRouter struct {
}

// InitJackpotRouter 初始化 Jackpot 路由信息
func (s *JackpotRouter) InitJackpotRouter(Router *gin.RouterGroup) {
	jackpotRouter := Router.Group("jackpot").Use(middleware.OperationRecord())
	jackpotRouterWithoutRecord := Router.Group("jackpot")
	var jackpotApi = v1.ApiGroupApp.BusinessApiGroup.JackpotApi
	{
		jackpotRouter.POST("createJackpot", jackpotApi.CreateJackpot)             // 新建Jackpot
		jackpotRouter.DELETE("deleteJackpot", jackpotApi.DeleteJackpot)           // 删除Jackpot
		jackpotRouter.DELETE("deleteJackpotByIds", jackpotApi.DeleteJackpotByIds) // 批量删除Jackpot
		jackpotRouter.PUT("updateJackpot", jackpotApi.UpdateJackpot)              // 更新Jackpot
		jackpotRouter.POST("saveJackpotList", jackpotApi.SaveJackpotList)         // 更新Jackpot
	}
	{
		jackpotRouterWithoutRecord.GET("findJackpot", jackpotApi.FindJackpot)       // 根据ID获取Jackpot
		jackpotRouterWithoutRecord.GET("getJackpotList", jackpotApi.GetJackpotList) // 获取Jackpot列表
	}
}
