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

type SlotSymbolApi struct {
}

var slotSymbolService = service.ServiceGroupApp.BusinessServiceGroup.SlotSymbolService

// CreateSlotSymbol 创建SlotSymbol
// @Tags SlotSymbol
// @Summary 创建SlotSymbol
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotSymbol true "创建SlotSymbol"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slotSymbol/createSlotSymbol [post]
func (slotSymbolApi *SlotSymbolApi) CreateSlotSymbol(c *gin.Context) {
	var slotSymbol business.SlotSymbol
	err := c.ShouldBindJSON(&slotSymbol)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"SlotId": {utils.NotEmpty()},
	}
	if err := utils.Verify(slotSymbol, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotSymbolService.CreateSlotSymbol(slotSymbol); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSlotSymbol 删除SlotSymbol
// @Tags SlotSymbol
// @Summary 删除SlotSymbol
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotSymbol true "删除SlotSymbol"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /slotSymbol/deleteSlotSymbol [delete]
func (slotSymbolApi *SlotSymbolApi) DeleteSlotSymbol(c *gin.Context) {
	var slotSymbol business.SlotSymbol
	err := c.ShouldBindJSON(&slotSymbol)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotSymbolService.DeleteSlotSymbol(slotSymbol); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSlotSymbolByIds 批量删除SlotSymbol
// @Tags SlotSymbol
// @Summary 批量删除SlotSymbol
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SlotSymbol"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /slotSymbol/deleteSlotSymbolByIds [delete]
func (slotSymbolApi *SlotSymbolApi) DeleteSlotSymbolByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotSymbolService.DeleteSlotSymbolByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSlotSymbol 更新SlotSymbol
// @Tags SlotSymbol
// @Summary 更新SlotSymbol
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotSymbol true "更新SlotSymbol"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /slotSymbol/updateSlotSymbol [put]
func (slotSymbolApi *SlotSymbolApi) UpdateSlotSymbol(c *gin.Context) {
	var slotSymbol business.SlotSymbol
	err := c.ShouldBindJSON(&slotSymbol)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"SlotId": {utils.NotEmpty()},
	}
	if err := utils.Verify(slotSymbol, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotSymbolService.UpdateSlotSymbol(slotSymbol); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSlotSymbol 用id查询SlotSymbol
// @Tags SlotSymbol
// @Summary 用id查询SlotSymbol
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.SlotSymbol true "用id查询SlotSymbol"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /slotSymbol/findSlotSymbol [get]
func (slotSymbolApi *SlotSymbolApi) FindSlotSymbol(c *gin.Context) {
	var slotSymbol business.SlotSymbol
	err := c.ShouldBindQuery(&slotSymbol)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reslotSymbol, err := slotSymbolService.GetSlotSymbol(slotSymbol.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败: "+err.Error(), c)
	} else {
		response.OkWithData(gin.H{"reslotSymbol": reslotSymbol}, c)
	}
}

// GetSlotSymbolList 分页获取SlotSymbol列表
// @Tags SlotSymbol
// @Summary 分页获取SlotSymbol列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.SlotSymbolSearch true "分页获取SlotSymbol列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slotSymbol/getSlotSymbolList [get]
func (slotSymbolApi *SlotSymbolApi) GetSlotSymbolList(c *gin.Context) {
	var pageInfo businessReq.SlotSymbolSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := slotSymbolService.GetSlotSymbolInfoList(pageInfo); err != nil {
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
