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

type SlotPayTableApi struct {
}

var slotPayTableService = service.ServiceGroupApp.BusinessServiceGroup.SlotPayTableService

// CreateSlotPayTable 创建SlotPayTable
// @Tags SlotPayTable
// @Summary 创建SlotPayTable
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotPayTable true "创建SlotPayTable"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slotPayTable/createSlotPayTable [post]
func (slotPayTableApi *SlotPayTableApi) CreateSlotPayTable(c *gin.Context) {
	var slotPayTable business.SlotPayTable
	err := c.ShouldBindJSON(&slotPayTable)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"SlotId": {utils.NotEmpty()},
		"Type":   {utils.NotEmpty()},
	}
	if err := utils.Verify(slotPayTable, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotPayTableService.CreateSlotPayTable(slotPayTable); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSlotPayTable 删除SlotPayTable
// @Tags SlotPayTable
// @Summary 删除SlotPayTable
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotPayTable true "删除SlotPayTable"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /slotPayTable/deleteSlotPayTable [delete]
func (slotPayTableApi *SlotPayTableApi) DeleteSlotPayTable(c *gin.Context) {
	var slotPayTable business.SlotPayTable
	err := c.ShouldBindJSON(&slotPayTable)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotPayTableService.DeleteSlotPayTable(slotPayTable); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSlotPayTableByIds 批量删除SlotPayTable
// @Tags SlotPayTable
// @Summary 批量删除SlotPayTable
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SlotPayTable"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /slotPayTable/deleteSlotPayTableByIds [delete]
func (slotPayTableApi *SlotPayTableApi) DeleteSlotPayTableByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotPayTableService.DeleteSlotPayTableByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSlotPayTable 更新SlotPayTable
// @Tags SlotPayTable
// @Summary 更新SlotPayTable
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotPayTable true "更新SlotPayTable"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /slotPayTable/updateSlotPayTable [put]
func (slotPayTableApi *SlotPayTableApi) UpdateSlotPayTable(c *gin.Context) {
	var slotPayTable business.SlotPayTable
	err := c.ShouldBindJSON(&slotPayTable)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"SlotId": {utils.NotEmpty()},
		"Type":   {utils.NotEmpty()},
	}
	if err := utils.Verify(slotPayTable, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotPayTableService.UpdateSlotPayTable(slotPayTable); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSlotPayTable 用id查询SlotPayTable
// @Tags SlotPayTable
// @Summary 用id查询SlotPayTable
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.SlotPayTable true "用id查询SlotPayTable"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /slotPayTable/findSlotPayTable [get]
func (slotPayTableApi *SlotPayTableApi) FindSlotPayTable(c *gin.Context) {
	var slotPayTable business.SlotPayTable
	err := c.ShouldBindQuery(&slotPayTable)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reslotPayTable, err := slotPayTableService.GetSlotPayTable(slotPayTable.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败: "+err.Error(), c)
	} else {
		response.OkWithData(gin.H{"reslotPayTable": reslotPayTable}, c)
	}
}

// GetSlotPayTableList 分页获取SlotPayTable列表
// @Tags SlotPayTable
// @Summary 分页获取SlotPayTable列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.SlotPayTableSearch true "分页获取SlotPayTable列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slotPayTable/getSlotPayTableList [get]
func (slotPayTableApi *SlotPayTableApi) GetSlotPayTableList(c *gin.Context) {
	var pageInfo businessReq.SlotPayTableSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := slotPayTableService.GetSlotPayTableInfoList(pageInfo); err != nil {
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
