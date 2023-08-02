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

type SlotReelApi struct {
}

var slotReelService = service.ServiceGroupApp.BusinessServiceGroup.SlotReelService

// CreateSlotReel 创建SlotReel
// @Tags SlotReel
// @Summary 创建SlotReel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotReel true "创建SlotReel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slotReel/createSlotReel [post]
func (slotReelApi *SlotReelApi) CreateSlotReel(c *gin.Context) {
	var slotReel business.SlotReel
	err := c.ShouldBindJSON(&slotReel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"SlotId": {utils.NotEmpty()},
	}
	if err := utils.Verify(slotReel, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotReelService.CreateSlotReel(slotReel); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSlotReel 删除SlotReel
// @Tags SlotReel
// @Summary 删除SlotReel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotReel true "删除SlotReel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /slotReel/deleteSlotReel [delete]
func (slotReelApi *SlotReelApi) DeleteSlotReel(c *gin.Context) {
	var slotReel business.SlotReel
	err := c.ShouldBindJSON(&slotReel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotReelService.DeleteSlotReel(slotReel); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSlotReelByIds 批量删除SlotReel
// @Tags SlotReel
// @Summary 批量删除SlotReel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SlotReel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /slotReel/deleteSlotReelByIds [delete]
func (slotReelApi *SlotReelApi) DeleteSlotReelByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotReelService.DeleteSlotReelByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSlotReel 更新SlotReel
// @Tags SlotReel
// @Summary 更新SlotReel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotReel true "更新SlotReel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /slotReel/updateSlotReel [put]
func (slotReelApi *SlotReelApi) UpdateSlotReel(c *gin.Context) {
	var slotReel business.SlotReel
	err := c.ShouldBindJSON(&slotReel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"SlotId": {utils.NotEmpty()},
	}
	if err := utils.Verify(slotReel, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotReelService.UpdateSlotReel(slotReel); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSlotReel 用id查询SlotReel
// @Tags SlotReel
// @Summary 用id查询SlotReel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.SlotReel true "用id查询SlotReel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /slotReel/findSlotReel [get]
func (slotReelApi *SlotReelApi) FindSlotReel(c *gin.Context) {
	var slotReel business.SlotReel
	err := c.ShouldBindQuery(&slotReel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reslotReel, err := slotReelService.GetSlotReel(slotReel.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败: "+err.Error(), c)
	} else {
		response.OkWithData(gin.H{"reslotReel": reslotReel}, c)
	}
}

// GetSlotReelList 分页获取SlotReel列表
// @Tags SlotReel
// @Summary 分页获取SlotReel列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.SlotReelSearch true "分页获取SlotReel列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slotReel/getSlotReelList [get]
func (slotReelApi *SlotReelApi) GetSlotReelList(c *gin.Context) {
	var pageInfo businessReq.SlotReelSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := slotReelService.GetSlotReelInfoList(pageInfo); err != nil {
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
