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

type SlotPaylineApi struct {
}

var slotPaylineService = service.ServiceGroupApp.BusinessServiceGroup.SlotPaylineService

// CreateSlotPayline 创建SlotPayline
// @Tags SlotPayline
// @Summary 创建SlotPayline
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotPayline true "创建SlotPayline"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slotPayline/createSlotPayline [post]
func (slotPaylineApi *SlotPaylineApi) CreateSlotPayline(c *gin.Context) {
	var slotPayline business.SlotPayline
	err := c.ShouldBindJSON(&slotPayline)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"No": {utils.NotEmpty()},
	}
	if err := utils.Verify(slotPayline, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotPaylineService.CreateSlotPayline(slotPayline); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSlotPayline 删除SlotPayline
// @Tags SlotPayline
// @Summary 删除SlotPayline
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotPayline true "删除SlotPayline"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /slotPayline/deleteSlotPayline [delete]
func (slotPaylineApi *SlotPaylineApi) DeleteSlotPayline(c *gin.Context) {
	var slotPayline business.SlotPayline
	err := c.ShouldBindJSON(&slotPayline)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotPaylineService.DeleteSlotPayline(slotPayline); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSlotPaylineByIds 批量删除SlotPayline
// @Tags SlotPayline
// @Summary 批量删除SlotPayline
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SlotPayline"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /slotPayline/deleteSlotPaylineByIds [delete]
func (slotPaylineApi *SlotPaylineApi) DeleteSlotPaylineByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotPaylineService.DeleteSlotPaylineByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSlotPayline 更新SlotPayline
// @Tags SlotPayline
// @Summary 更新SlotPayline
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotPayline true "更新SlotPayline"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /slotPayline/updateSlotPayline [put]
func (slotPaylineApi *SlotPaylineApi) UpdateSlotPayline(c *gin.Context) {
	var slotPayline business.SlotPayline
	err := c.ShouldBindJSON(&slotPayline)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"No": {utils.NotEmpty()},
	}
	if err := utils.Verify(slotPayline, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotPaylineService.UpdateSlotPayline(slotPayline); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSlotPayline 用id查询SlotPayline
// @Tags SlotPayline
// @Summary 用id查询SlotPayline
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.SlotPayline true "用id查询SlotPayline"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /slotPayline/findSlotPayline [get]
func (slotPaylineApi *SlotPaylineApi) FindSlotPayline(c *gin.Context) {
	var slotPayline business.SlotPayline
	err := c.ShouldBindQuery(&slotPayline)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reslotPayline, err := slotPaylineService.GetSlotPayline(slotPayline.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败: "+err.Error(), c)
	} else {
		response.OkWithData(gin.H{"reslotPayline": reslotPayline}, c)
	}
}

// GetSlotPaylineList 分页获取SlotPayline列表
// @Tags SlotPayline
// @Summary 分页获取SlotPayline列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.SlotPaylineSearch true "分页获取SlotPayline列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slotPayline/getSlotPaylineList [get]
func (slotPaylineApi *SlotPaylineApi) GetSlotPaylineList(c *gin.Context) {
	var pageInfo businessReq.SlotPaylineSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := slotPaylineService.GetSlotPaylineInfoList(pageInfo); err != nil {
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
