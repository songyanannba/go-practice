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
	businessService "slot-server/service/business"
)

type SlotApi struct {
}

var slotService = service.ServiceGroupApp.BusinessServiceGroup.SlotService

// CreateSlot 创建Slot
// @Tags Slot
// @Summary 创建Slot
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.Slot true "创建Slot"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slot/createSlot [post]
func (slotApi *SlotApi) CreateSlot(c *gin.Context) {
	var slot business.Slot
	err := c.ShouldBindJSON(&slot)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotService.CreateSlot(slot); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSlot 删除Slot
// @Tags Slot
// @Summary 删除Slot
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.Slot true "删除Slot"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /slot/deleteSlot [delete]
func (slotApi *SlotApi) DeleteSlot(c *gin.Context) {
	var slot business.Slot
	err := c.ShouldBindJSON(&slot)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotService.DeleteSlot(slot); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSlotByIds 批量删除Slot
// @Tags Slot
// @Summary 批量删除Slot
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Slot"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /slot/deleteSlotByIds [delete]
func (slotApi *SlotApi) DeleteSlotByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotService.DeleteSlotByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSlot 更新Slot
// @Tags Slot
// @Summary 更新Slot
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.Slot true "更新Slot"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /slot/updateSlot [put]
func (slotApi *SlotApi) UpdateSlot(c *gin.Context) {
	var slot business.Slot
	err := c.ShouldBindJSON(&slot)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotService.UpdateSlot(slot); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSlot 用id查询Slot
// @Tags Slot
// @Summary 用id查询Slot
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.Slot true "用id查询Slot"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /slot/findSlot [get]
func (slotApi *SlotApi) FindSlot(c *gin.Context) {
	var slot business.Slot
	err := c.ShouldBindQuery(&slot)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reslot, err := slotService.GetSlot(slot.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败: "+err.Error(), c)
	} else {
		response.OkWithData(gin.H{"reslot": reslot}, c)
	}
}

// GetSlotList 分页获取Slot列表
// @Tags Slot
// @Summary 分页获取Slot列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.SlotSearch true "分页获取Slot列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slot/getSlotList [get]
func (slotApi *SlotApi) GetSlotList(c *gin.Context) {
	var pageInfo businessReq.SlotSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := slotService.GetSlotInfoList(pageInfo); err != nil {
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

func (slotApi *SlotApi) UpdateAllConfig(c *gin.Context) {
	var (
		params businessService.AllConfigParams
	)
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	text, success := slotService.UpdateAllConfig(&params)
	if !success {
		global.GVA_LOG.Error("更新所有配置失败!", zap.String("err", text))
		response.FailWithMessage(text, c)
		return
	}
	response.OkWithMessage(text, c)
}

func (slotApi *SlotApi) BackendOperate(c *gin.Context) {
	var (
		params businessService.BackendOperateParams
		msg    string
	)
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	token := c.Request.Header.Get("x-token")
	msg, err = slotService.BackendOperate(token, &params)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	response.OkWithMessage(msg, c)
}
