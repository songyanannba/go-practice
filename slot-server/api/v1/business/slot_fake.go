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

type SlotFakeApi struct {
}

var slotFakeService = service.ServiceGroupApp.BusinessServiceGroup.SlotFakeService

// CreateSlotFake 创建SlotFake
// @Tags SlotFake
// @Summary 创建SlotFake
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotFake true "创建SlotFake"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slotFake/createSlotFake [post]
func (slotFakeApi *SlotFakeApi) CreateSlotFake(c *gin.Context) {
	var slotFake business.SlotFake
	err := c.ShouldBindJSON(&slotFake)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotFakeService.CreateSlotFake(slotFake); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSlotFake 删除SlotFake
// @Tags SlotFake
// @Summary 删除SlotFake
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotFake true "删除SlotFake"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /slotFake/deleteSlotFake [delete]
func (slotFakeApi *SlotFakeApi) DeleteSlotFake(c *gin.Context) {
	var slotFake business.SlotFake
	err := c.ShouldBindJSON(&slotFake)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotFakeService.DeleteSlotFake(slotFake); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSlotFakeByIds 批量删除SlotFake
// @Tags SlotFake
// @Summary 批量删除SlotFake
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SlotFake"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /slotFake/deleteSlotFakeByIds [delete]
func (slotFakeApi *SlotFakeApi) DeleteSlotFakeByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotFakeService.DeleteSlotFakeByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSlotFake 更新SlotFake
// @Tags SlotFake
// @Summary 更新SlotFake
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotFake true "更新SlotFake"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /slotFake/updateSlotFake [put]
func (slotFakeApi *SlotFakeApi) UpdateSlotFake(c *gin.Context) {
	var slotFake business.SlotFake
	err := c.ShouldBindJSON(&slotFake)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotFakeService.UpdateSlotFake(slotFake); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSlotFake 用id查询SlotFake
// @Tags SlotFake
// @Summary 用id查询SlotFake
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.SlotFake true "用id查询SlotFake"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /slotFake/findSlotFake [get]
func (slotFakeApi *SlotFakeApi) FindSlotFake(c *gin.Context) {
	var slotFake business.SlotFake
	err := c.ShouldBindQuery(&slotFake)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reslotFake, err := slotFakeService.GetSlotFake(slotFake.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败: "+err.Error(), c)
	} else {
		response.OkWithData(gin.H{"reslotFake": reslotFake}, c)
	}
}

// GetSlotFakeList 分页获取SlotFake列表
// @Tags SlotFake
// @Summary 分页获取SlotFake列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.SlotFakeSearch true "分页获取SlotFake列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slotFake/getSlotFakeList [get]
func (slotFakeApi *SlotFakeApi) GetSlotFakeList(c *gin.Context) {
	var pageInfo businessReq.SlotFakeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := slotFakeService.GetSlotFakeInfoList(pageInfo); err != nil {
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
