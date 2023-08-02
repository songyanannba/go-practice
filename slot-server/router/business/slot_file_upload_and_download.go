package business

import (
	"github.com/gin-gonic/gin"
	"slot-server/api/v1"
	"slot-server/middleware"
)

type SlotFileUploadAndDownloadRouter struct {
}

// InitSlotFileUploadAndDownloadRouter 初始化 SlotFileUploadAndDownload 路由信息
func (s *SlotFileUploadAndDownloadRouter) InitSlotFileUploadAndDownloadRouter(Router *gin.RouterGroup) {
	SlotFileUpAndDownRouter := Router.Group("SlotFileUpAndDown").Use(middleware.OperationRecord())
	SlotFileUpAndDownRouterWithoutRecord := Router.Group("SlotFileUpAndDown")
	var SlotFileUpAndDownApi = v1.ApiGroupApp.BusinessApiGroup.SlotFileUploadAndDownloadApi
	{
		SlotFileUpAndDownRouter.POST("upload", SlotFileUpAndDownApi.UploadFile)
		SlotFileUpAndDownRouter.POST("createSlotFileUploadAndDownload", SlotFileUpAndDownApi.CreateSlotFileUploadAndDownload)   // 新建SlotFileUploadAndDownload
		SlotFileUpAndDownRouter.DELETE("deleteSlotFileUploadAndDownload", SlotFileUpAndDownApi.DeleteSlotFileUploadAndDownload) // 删除SlotFileUploadAndDownload
		SlotFileUpAndDownRouter.DELETE("deleteSlotFileUploadAndDownloadByIds", SlotFileUpAndDownApi.DeleteSlotFileUploadAndDownloadByIds) // 批量删除SlotFileUploadAndDownload
		SlotFileUpAndDownRouter.PUT("updateSlotFileUploadAndDownload", SlotFileUpAndDownApi.UpdateSlotFileUploadAndDownload)    // 更新SlotFileUploadAndDownload
	}
	{
		SlotFileUpAndDownRouterWithoutRecord.GET("findSlotFileUploadAndDownload", SlotFileUpAndDownApi.FindSlotFileUploadAndDownload)        // 根据ID获取SlotFileUploadAndDownload
		SlotFileUpAndDownRouterWithoutRecord.GET("getSlotFileUploadAndDownloadList", SlotFileUpAndDownApi.GetSlotFileUploadAndDownloadList)  // 获取SlotFileUploadAndDownload列表
	}
}
