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
)

type SlotUserSpinApi struct {
}

var slotUserSpinService = service.ServiceGroupApp.BusinessServiceGroup.SlotUserSpinService

// CreateSlotUserSpin 创建SlotUserSpin
// @Tags SlotUserSpin
// @Summary 创建SlotUserSpin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotUserSpin true "创建SlotUserSpin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slotUserSpin/createSlotUserSpin [post]
func (slotUserSpinApi *SlotUserSpinApi) CreateSlotUserSpin(c *gin.Context) {
	var slotUserSpin business.SlotUserSpin
	err := c.ShouldBindJSON(&slotUserSpin)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotUserSpinService.CreateSlotUserSpin(slotUserSpin); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSlotUserSpin 删除SlotUserSpin
// @Tags SlotUserSpin
// @Summary 删除SlotUserSpin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotUserSpin true "删除SlotUserSpin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /slotUserSpin/deleteSlotUserSpin [delete]
func (slotUserSpinApi *SlotUserSpinApi) DeleteSlotUserSpin(c *gin.Context) {
	var slotUserSpin business.SlotUserSpin
	err := c.ShouldBindJSON(&slotUserSpin)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotUserSpinService.DeleteSlotUserSpin(slotUserSpin); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSlotUserSpinByIds 批量删除SlotUserSpin
// @Tags SlotUserSpin
// @Summary 批量删除SlotUserSpin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SlotUserSpin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /slotUserSpin/deleteSlotUserSpinByIds [delete]
func (slotUserSpinApi *SlotUserSpinApi) DeleteSlotUserSpinByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotUserSpinService.DeleteSlotUserSpinByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSlotUserSpin 更新SlotUserSpin
// @Tags SlotUserSpin
// @Summary 更新SlotUserSpin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotUserSpin true "更新SlotUserSpin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /slotUserSpin/updateSlotUserSpin [put]
func (slotUserSpinApi *SlotUserSpinApi) UpdateSlotUserSpin(c *gin.Context) {
	var slotUserSpin business.SlotUserSpin
	err := c.ShouldBindJSON(&slotUserSpin)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotUserSpinService.UpdateSlotUserSpin(slotUserSpin); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSlotUserSpin 用id查询SlotUserSpin
// @Tags SlotUserSpin
// @Summary 用id查询SlotUserSpin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.SlotUserSpin true "用id查询SlotUserSpin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /slotUserSpin/findSlotUserSpin [get]
func (slotUserSpinApi *SlotUserSpinApi) FindSlotUserSpin(c *gin.Context) {
	var slotUserSpin business.SlotUserSpin
	err := c.ShouldBindQuery(&slotUserSpin)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reslotUserSpin, err := slotUserSpinService.GetSlotUserSpin(slotUserSpin.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败: "+err.Error(), c)
	} else {
		response.OkWithData(gin.H{"reslotUserSpin": reslotUserSpin}, c)
	}
}

// GetSlotUserSpinList 分页获取SlotUserSpin列表
// @Tags SlotUserSpin
// @Summary 分页获取SlotUserSpin列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.SlotUserSpinSearch true "分页获取SlotUserSpin列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slotUserSpin/getSlotUserSpinList [get]
func (slotUserSpinApi *SlotUserSpinApi) GetSlotUserSpinList(c *gin.Context) {
	var pageInfo businessReq.SlotUserSpinSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := slotUserSpinService.GetSlotUserSpinInfoList(pageInfo); err != nil {
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
