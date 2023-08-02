package business

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
	"slot-server/model/common/response"
	"slot-server/service"
	"slot-server/utils"
)

type SlotReelDataApi struct {
}

var slotReelDataService = service.ServiceGroupApp.BusinessServiceGroup.SlotReelDataService

// CreateSlotReelData 创建SlotReelData
// @Tags SlotReelData
// @Summary 创建SlotReelData
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotReelData true "创建SlotReelData"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slotReelData/createSlotReelData [post]
func (slotReelDataApi *SlotReelDataApi) CreateSlotReelData(c *gin.Context) {
	var slotReelData business.SlotReelData
	err := c.ShouldBindJSON(&slotReelData)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"SlotId": {utils.NotEmpty()},
	}
	if err := utils.Verify(slotReelData, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotReelDataService.CreateSlotReelData(slotReelData); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSlotReelData 删除SlotReelData
// @Tags SlotReelData
// @Summary 删除SlotReelData
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotReelData true "删除SlotReelData"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /slotReelData/deleteSlotReelData [delete]
func (slotReelDataApi *SlotReelDataApi) DeleteSlotReelData(c *gin.Context) {
	var slotReelData business.SlotReelData
	err := c.ShouldBindJSON(&slotReelData)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotReelDataService.DeleteSlotReelData(slotReelData); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSlotReelDataByIds 批量删除SlotReelData
// @Tags SlotReelData
// @Summary 批量删除SlotReelData
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SlotReelData"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /slotReelData/deleteSlotReelDataByIds [delete]
func (slotReelDataApi *SlotReelDataApi) DeleteSlotReelDataByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotReelDataService.DeleteSlotReelDataByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSlotReelData 更新SlotReelData
// @Tags SlotReelData
// @Summary 更新SlotReelData
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotReelData true "更新SlotReelData"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /slotReelData/updateSlotReelData [put]
func (slotReelDataApi *SlotReelDataApi) UpdateSlotReelData(c *gin.Context) {
	var slotReelData business.SlotReelData
	err := c.ShouldBindJSON(&slotReelData)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"SlotId": {utils.NotEmpty()},
	}
	if err := utils.Verify(slotReelData, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotReelDataService.UpdateSlotReelData(slotReelData); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSlotReelData 用id查询SlotReelData
// @Tags SlotReelData
// @Summary 用id查询SlotReelData
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.SlotReelData true "用id查询SlotReelData"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /slotReelData/findSlotReelData [get]
func (slotReelDataApi *SlotReelDataApi) FindSlotReelData(c *gin.Context) {
	var slotReelData business.SlotReelData
	err := c.ShouldBindQuery(&slotReelData)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reslotReelData, err := slotReelDataService.GetSlotReelData(slotReelData.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败: "+err.Error(), c)
	} else {
		response.OkWithData(gin.H{"reslotReelData": reslotReelData}, c)
	}
}

// GetSlotReelDataList 分页获取SlotReelData列表
// @Tags SlotReelData
// @Summary 分页获取SlotReelData列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.SlotReelDataSearch true "分页获取SlotReelData列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slotReelData/getSlotReelDataList [get]
func (slotReelDataApi *SlotReelDataApi) GetSlotReelDataList(c *gin.Context) {
	var pageInfo businessReq.SlotReelDataSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := slotReelDataService.GetSlotReelDataInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败: "+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
