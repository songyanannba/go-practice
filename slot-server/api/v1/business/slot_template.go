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

type SlotTemplateApi struct {
}

var slotTemplateService = service.ServiceGroupApp.BusinessServiceGroup.SlotTemplateService

// CreateSlotTemplate 创建SlotTemplate
// @Tags SlotTemplate
// @Summary 创建SlotTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotTemplate true "创建SlotTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slotTemplate/createSlotTemplate [post]
func (slotTemplateApi *SlotTemplateApi) CreateSlotTemplate(c *gin.Context) {
	var slotTemplate business.SlotTemplate
	err := c.ShouldBindJSON(&slotTemplate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"SlotId": {utils.NotEmpty()},
		"Type":   {utils.NotEmpty()},
		//"Column": {utils.NotEmpty()},
		"Layout": {utils.NotEmpty()},
	}
	if err := utils.Verify(slotTemplate, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotTemplateService.CreateSlotTemplate(slotTemplate); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSlotTemplate 删除SlotTemplate
// @Tags SlotTemplate
// @Summary 删除SlotTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotTemplate true "删除SlotTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /slotTemplate/deleteSlotTemplate [delete]
func (slotTemplateApi *SlotTemplateApi) DeleteSlotTemplate(c *gin.Context) {
	var slotTemplate business.SlotTemplate
	err := c.ShouldBindJSON(&slotTemplate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotTemplateService.DeleteSlotTemplate(slotTemplate); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSlotTemplateByIds 批量删除SlotTemplate
// @Tags SlotTemplate
// @Summary 批量删除SlotTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SlotTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /slotTemplate/deleteSlotTemplateByIds [delete]
func (slotTemplateApi *SlotTemplateApi) DeleteSlotTemplateByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotTemplateService.DeleteSlotTemplateByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSlotTemplate 更新SlotTemplate
// @Tags SlotTemplate
// @Summary 更新SlotTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotTemplate true "更新SlotTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /slotTemplate/updateSlotTemplate [put]
func (slotTemplateApi *SlotTemplateApi) UpdateSlotTemplate(c *gin.Context) {
	var slotTemplate business.SlotTemplate
	err := c.ShouldBindJSON(&slotTemplate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"SlotId": {utils.NotEmpty()},
		"Type":   {utils.NotEmpty()},
		"Column": {utils.NotEmpty()},
		"Layout": {utils.NotEmpty()},
	}
	if err := utils.Verify(slotTemplate, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotTemplateService.UpdateSlotTemplate(slotTemplate); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSlotTemplate 用id查询SlotTemplate
// @Tags SlotTemplate
// @Summary 用id查询SlotTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.SlotTemplate true "用id查询SlotTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /slotTemplate/findSlotTemplate [get]
func (slotTemplateApi *SlotTemplateApi) FindSlotTemplate(c *gin.Context) {
	var slotTemplate business.SlotTemplate
	err := c.ShouldBindQuery(&slotTemplate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reslotTemplate, err := slotTemplateService.GetSlotTemplate(slotTemplate.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败: "+err.Error(), c)
	} else {
		response.OkWithData(gin.H{"reslotTemplate": reslotTemplate}, c)
	}
}

// GetSlotTemplateList 分页获取SlotTemplate列表
// @Tags SlotTemplate
// @Summary 分页获取SlotTemplate列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.SlotTemplateSearch true "分页获取SlotTemplate列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slotTemplate/getSlotTemplateList [get]
func (slotTemplateApi *SlotTemplateApi) GetSlotTemplateList(c *gin.Context) {
	var pageInfo businessReq.SlotTemplateSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := slotTemplateService.GetSlotTemplateInfoList(pageInfo); err != nil {
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
