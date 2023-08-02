package business

import (
	"github.com/gin-gonic/gin"
	"slot-server/api/v1"
	"slot-server/middleware"
)

type SlotReelRouter struct {
}

// InitSlotReelRouter 初始化 SlotReel 路由信息
func (s *SlotReelRouter) InitSlotReelRouter(Router *gin.RouterGroup) {
	slotReelRouter := Router.Group("slotReel").Use(middleware.OperationRecord())
	slotReelRouterWithoutRecord := Router.Group("slotReel")
	var slotReelApi = v1.ApiGroupApp.BusinessApiGroup.SlotReelApi
	{
		slotReelRouter.POST("createSlotReel", slotReelApi.CreateSlotReel)             // 新建SlotReel
		slotReelRouter.DELETE("deleteSlotReel", slotReelApi.DeleteSlotReel)           // 删除SlotReel
		slotReelRouter.DELETE("deleteSlotReelByIds", slotReelApi.DeleteSlotReelByIds) // 批量删除SlotReel
		slotReelRouter.PUT("updateSlotReel", slotReelApi.UpdateSlotReel)              // 更新SlotReel
	}
	{
		slotReelRouterWithoutRecord.GET("findSlotReel", slotReelApi.FindSlotReel)       // 根据ID获取SlotReel
		slotReelRouterWithoutRecord.GET("getSlotReelList", slotReelApi.GetSlotReelList) // 获取SlotReel列表
	}
}
