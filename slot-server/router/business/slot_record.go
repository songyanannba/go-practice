package business

import (
	"github.com/gin-gonic/gin"
	"slot-server/api/v1"
	"slot-server/middleware"
)

type SlotRecordRouter struct {
}

// InitSlotRecordRouter 初始化 SlotRecord 路由信息
func (s *SlotRecordRouter) InitSlotRecordRouter(Router *gin.RouterGroup) {
	RecordRouter := Router.Group("Record").Use(middleware.OperationRecord())
	RecordRouterWithoutRecord := Router.Group("Record")
	var RecordApi = v1.ApiGroupApp.BusinessApiGroup.SlotRecordApi
	{
		RecordRouter.POST("createSlotRecord", RecordApi.CreateSlotRecord)             // 新建SlotRecord
		RecordRouter.DELETE("deleteSlotRecord", RecordApi.DeleteSlotRecord)           // 删除SlotRecord
		RecordRouter.DELETE("deleteSlotRecordByIds", RecordApi.DeleteSlotRecordByIds) // 批量删除SlotRecord
		RecordRouter.PUT("updateSlotRecord", RecordApi.UpdateSlotRecord)              // 更新SlotRecord
	}
	{
		RecordRouterWithoutRecord.GET("findSlotRecord", RecordApi.FindSlotRecord)       // 根据ID获取SlotRecord
		RecordRouterWithoutRecord.GET("getSlotRecordList", RecordApi.GetSlotRecordList) // 获取SlotRecord列表
	}
}
