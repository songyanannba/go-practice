package business

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
	"slot-server/model/common/response"
	"slot-server/model/system"
	"slot-server/service"
	"slot-server/utils"
)

type SlotRecordApi struct {
}

var RecordService = service.ServiceGroupApp.BusinessServiceGroup.SlotRecordService

// CreateSlotRecord 创建SlotRecord
// @Tags SlotRecord
// @Summary 创建SlotRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotRecord true "创建SlotRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Record/createSlotRecord [post]
func (RecordApi *SlotRecordApi) CreateSlotRecord(c *gin.Context) {
	var Record business.SlotRecord
	err := c.ShouldBindJSON(&Record)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"UserId": {utils.NotEmpty()},
		"slotId": {utils.NotEmpty()},
	}
	if err := utils.Verify(Record, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := RecordService.CreateSlotRecord(Record); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSlotRecord 删除SlotRecord
// @Tags SlotRecord
// @Summary 删除SlotRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotRecord true "删除SlotRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Record/deleteSlotRecord [delete]
func (RecordApi *SlotRecordApi) DeleteSlotRecord(c *gin.Context) {
	var Record business.SlotRecord
	err := c.ShouldBindJSON(&Record)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := RecordService.DeleteSlotRecord(Record); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSlotRecordByIds 批量删除SlotRecord
// @Tags SlotRecord
// @Summary 批量删除SlotRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SlotRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Record/deleteSlotRecordByIds [delete]
func (RecordApi *SlotRecordApi) DeleteSlotRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := RecordService.DeleteSlotRecordByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSlotRecord 更新SlotRecord
// @Tags SlotRecord
// @Summary 更新SlotRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotRecord true "更新SlotRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Record/updateSlotRecord [put]
func (RecordApi *SlotRecordApi) UpdateSlotRecord(c *gin.Context) {
	var Record business.SlotRecord
	err := c.ShouldBindJSON(&Record)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"UserId": {utils.NotEmpty()},
		"slotId": {utils.NotEmpty()},
	}
	if err := utils.Verify(Record, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := RecordService.UpdateSlotRecord(Record); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSlotRecord 用id查询SlotRecord
// @Tags SlotRecord
// @Summary 用id查询SlotRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.SlotRecord true "用id查询SlotRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Record/findSlotRecord [get]
func (RecordApi *SlotRecordApi) FindSlotRecord(c *gin.Context) {
	var Record business.SlotRecord
	err := c.ShouldBindQuery(&Record)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reRecord, err := RecordService.GetSlotRecord(Record.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败: "+err.Error(), c)
	} else {
		response.OkWithData(gin.H{"reRecord": reRecord}, c)
	}
}

// GetSlotRecordList 分页获取SlotRecord列表
// @Tags SlotRecord
// @Summary 分页获取SlotRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.SlotRecordSearch true "分页获取SlotRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Record/getSlotRecordList [get]
func (RecordApi *SlotRecordApi) GetSlotRecordList(c *gin.Context) {
	var pageInfo businessReq.SlotRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	info := utils.GetUserInfo(c)
	if info.AuthorityId == 10 {
		err = global.GVA_DB.Model(&system.SysUser{}).
			Where("id = ?", info.ID).
			Pluck("merchant_id", &pageInfo.MerchantId).Error
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
	}
	if list, total, err := RecordService.GetSlotRecordInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败: "+err.Error(), c)
	} else {
		response.OkWithDetailed(map[string]any{
			"list":     list,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
			"domain":   global.GVA_CONFIG.System.GameDomain,
		}, "获取成功", c)
	}
}
